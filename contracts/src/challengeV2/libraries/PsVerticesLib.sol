// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "../DataEntities.sol";
import "./ChallengeVertexLib.sol";

/// @title  Presumptive Successor Vertices library
/// @notice A collection of challenge vertices linked by: predecessorId, psId and lowestHeightSuccessorId
///         This library allows vertices to be connected and these ids updated only in ways that preserve
///         presumptive successor behaviour.
///         A presumptive successor is a vertex with the lowest unique height above its predecessor. If other
///         vertices are tied with lowest height, then none of them are the presumptive successor.
///         Some invariants:
///             * A vertex cannot become presumptive after being non-presumptive without changing it's predecessor
///             * Once the presumptive successor timer of a vertex is a above the lock threshold it cannot be made non-presumptive
///             * A vertex which is ps must be the lowest height successor of its predecessor
///             * There can be no other vertices that have the same height and predecessor as a presumptive successor
///             * If a vertex has a presumptive successor, it is also the lowest height successor
///             * If a vertex has no lowest height successor (so no successor at all), then it cannot have a presumptive successor
///             * The lowest height successor of a vertex can only decrease, never increase
///             * There is a ps threshold, once a successor has a ps timer greater than the threshold then they will remain ps forever
///             * New successor cannot be connected to a vertex whose ps has a timer greater than the threshold
library PsVerticesLib {
    using ChallengeVertexLib for ChallengeVertex;

    /// @notice Check that the vertex is the root of a one step fork. A one step fork is where 2 or more
    ///         vertices are successors to this vertex, and have a height exactly one greater than the height of this vertex.
    /// @param vertices The vertices collection
    /// @param vId      The one step fork root to check
    function checkAtOneStepFork(mapping(bytes32 => ChallengeVertex) storage vertices, bytes32 vId) internal view {
        require(vertices[vId].exists(), "Fork candidate vertex does not exist");
        require(!vertices[vId].isLeaf(), "Leaf can never be a fork candidate");

        // if this vertex has no successor at all, it cannot be the root of a one step fork
        require(vertices[vertices[vId].lowestHeightSuccessorId].exists(), "No successors");

        // the lowest height must be the root height + 1 at a one step fork
        uint256 lowestHeightSuccessorHeight = vertices[vertices[vId].lowestHeightSuccessorId].height;
        require(
            lowestHeightSuccessorHeight - vertices[vId].height == 1, "Lowest height not one above the current height"
        );

        // if 2 ore more successors are at the lowest height then the presumptive successor id is 0
        // therefore if the lowest height is 1 greater, and the presumptive successor id is 0 then we have
        // 2 or more successors at a height 1 greater than the root - so the root is a one step fork
        require(vertices[vId].psId == 0, "Has presumptive successor");
    }

    /// @notice Does the presumptive successor of the supplied vertex have a ps timer greater than the provided threshold
    /// @param vertices         The vertices collection
    /// @param vId              The vertex whose presumptive successor we are checking
    /// @param psThresholdSec   The ps threshold in seconds. A ps vertex cannot be made non presumptive if its timer exceeds the ps threshold
    function psExceedsPsThreshold(
        mapping(bytes32 => ChallengeVertex) storage vertices,
        bytes32 vId,
        uint256 psThresholdSec
    ) internal view returns (bool) {
        require(vertices[vId].exists(), "Predecessor vertex does not exist");

        // we dont allow presumptive successor to be updated if the ps has a timer that exceeds the ps threshold
        // therefore if it is at 0 must non of the successors must have a high enough timer,
        // or this is a new vertex so it doesnt have any successors, and therefore no high enough ps
        if (vertices[vId].psId == 0) {
            return false;
        }

        return getCurrentPsTimer(vertices, vertices[vId].psId) > psThresholdSec;
    }

    /// @notice The amount of time (seconds) this vertex has spent as the presumptive successor.
    ///         Use this function instead of the flushPsTime since this function also takes into account unflushed time
    /// @dev    We record ps time using the psLastUpdatedTimestamp on the predecessor vertex, and flush it onto the target it vertex
    ///         This means that the flushPsTime does not represent the total ps time where the vertex in question is currently the ps
    /// @param vertices The collection of vertices
    /// @param vId      The vertex whose ps timer we want to get
    function getCurrentPsTimer(mapping(bytes32 => ChallengeVertex) storage vertices, bytes32 vId)
        internal
        view
        returns (uint256)
    {
        require(vertices[vId].exists(), "Vertex does not exist for ps timer");
        bytes32 predecessorId = vertices[vId].predecessorId;
        require(vertices[predecessorId].exists(), "Predecessor vertex does not exist");

        if (vertices[predecessorId].psId == vId) {
            // if the vertex is currently the presumptive one we add the flushed time and the unflushed time
            return (block.timestamp - vertices[predecessorId].psLastUpdatedTimestamp) + vertices[vId].flushedPsTimeSec;
        } else {
            return vertices[vId].flushedPsTimeSec;
        }
    }

    /// @notice Flush the psLastUpdatedTimestamp of a vertex onto the current ps, and record that this occurred.
    ///         Once flushed will also check that the final flushed time is at least the provided minimum
    /// @param vertices             The ps vertices
    /// @param vId                  The id of the vertex on which to update psLastUpdatedTimestamp
    function flushPs(mapping(bytes32 => ChallengeVertex) storage vertices, bytes32 vId) internal {
        require(vertices[vId].exists(), "Vertex does not exist");
        // leaves should never have a ps, so we cant flush here
        require(!vertices[vId].isLeaf(), "Cannot flush leaf as it will never have a PS");

        // if a presumptive successor already exists we flush it
        if (vertices[vId].psId != 0) {
            uint256 timeToAdd = block.timestamp - vertices[vId].psLastUpdatedTimestamp;
            uint256 timeToSet = vertices[vertices[vId].psId].flushedPsTimeSec + timeToAdd;
            vertices[vertices[vId].psId].setFlushedPsTimeSec(timeToSet);
        }
        // every time we update the ps we record when it happened so that we can flush in the future
        vertices[vId].setPsLastUpdatedTimestamp(block.timestamp);
    }

    /// @notice Override the flushed ps time of a vertex.
    /// @dev    Does not allow an override above the threshold, this ensures that the threshold
    ///         can only be crossed by flushing, not by overriding
    /// @param vertices             The ps vertices
    /// @param vId                  The vertex to override the flushed ps time on
    /// @param newFlushedPsTimeSec  The new flushed ps time to set, must be less than psThresholdSec
    /// @param psThresholdSec       The ps threshold
    function overrideFlushedPsTime(
        mapping(bytes32 => ChallengeVertex) storage vertices,
        bytes32 vId,
        uint256 newFlushedPsTimeSec,
        uint256 psThresholdSec
    ) internal {
        require(vertices[vId].exists(), "Vertex does not exist");
        require(!vertices[vId].isRoot(), "Root must always have zero flushed ps time");

        // we dont allow overriding to cross the ps threshold - this ensures that
        // that the threshold can only be crossed by flushing
        require(newFlushedPsTimeSec < psThresholdSec, "Override crossed threshold");

        vertices[vId].setFlushedPsTimeSec(newFlushedPsTimeSec);
    }

    /// @notice Connect two existing vertices. The connection is made by setting the predecessor of the end vertex to
    ///         be the start vertex. When the connection is made ps timers, and lowest heigh successor, are updated
    ///         if relevant.
    /// @param vertices         The collection of vertices
    /// @param startVertexId    The start vertex to connect to
    /// @param endVertexId      The end vertex to connect from
    /// @param psThresholdSec   The ps threshold in seconds. A ps vertex cannot be made non presumptive if its timer exceeds the ps threshold
    function connect(
        mapping(bytes32 => ChallengeVertex) storage vertices,
        bytes32 startVertexId,
        bytes32 endVertexId,
        uint256 psThresholdSec
    ) internal {
        require(vertices[startVertexId].exists(), "Start vertex does not exist");
        // by definition of a leaf no connection can occur if the leaf is a start vertex
        require(!vertices[startVertexId].isLeaf(), "Cannot connect a successor to a leaf");
        require(vertices[endVertexId].exists(), "End vertex does not exist");
        require(vertices[endVertexId].predecessorId != startVertexId, "Vertices already connected");
        require(vertices[startVertexId].height < vertices[endVertexId].height, "Start height not lower than end height");
        // cannot connect vertices that are in different challenges
        require(
            vertices[startVertexId].challengeId == vertices[endVertexId].challengeId,
            "Predecessor and successor are in different challenges"
        );

        // always flush the current start vertex
        flushPs(vertices, startVertexId);

        // the start vertex has a ps that exceeds the threshold
        // we dont allow it to be connected to anything new
        require(
            !psExceedsPsThreshold(vertices, startVertexId, psThresholdSec),
            "Start vertex has ps with timer greater than ps threshold, cannot connect"
        );

        // now make the connection to the new predecessor
        vertices[endVertexId].setPredecessor(startVertexId);

        // the current vertex has no successors, in this case the new successor will certainly
        // be the ps
        bytes32 lowestHeightSuccessorId = vertices[startVertexId].lowestHeightSuccessorId;
        if (lowestHeightSuccessorId == 0) {
            // we flush here to update the timestamp on the start vertex
            vertices[startVertexId].setPsId(endVertexId);
            return;
        }

        uint256 height = vertices[endVertexId].height;
        uint256 lowestHeightSuccessorHeight = vertices[lowestHeightSuccessorId].height;
        // we're connect a successor that is lower than the current lowest height, so this new successor
        // will become the ps. Set the ps.
        if (height < lowestHeightSuccessorHeight) {
            vertices[startVertexId].setPsId(endVertexId);
            return;
        }

        // we're connecting a sibling to the current lowest height, that means that there will be more than
        // one successor at the same lowest height, in this case we set non of the successors to be the ps
        if (height == lowestHeightSuccessorHeight) {
            vertices[startVertexId].setPsId(0);
            return;
        }

        // if we're here we're connecting at a height above the lowest height successor
        // we dont need to make an ps updates
    }

    /// @notice Adds a vertex to the collection, and connects it to the provided predecessor
    /// @param vertices         The vertex collection
    /// @param vertex           The vertex to add
    /// @param predecessorId    The predecessor this vertex will become a successor to
    /// @param psThresholdSec   The ps threshold in seconds. A ps vertex cannot be made non presumptive if its timer exceeds the ps threshold
    function addVertex(
        mapping(bytes32 => ChallengeVertex) storage vertices,
        ChallengeVertex memory vertex,
        bytes32 predecessorId,
        uint256 psThresholdSec
    ) internal returns (bytes32) {
        bytes32 vId = vertex.id();
        require(!vertices[vId].exists(), "Vertex already exists");
        vertices[vId] = vertex;

        // connect the newly stored vertex to an existing vertex
        connect(vertices, predecessorId, vId, psThresholdSec);

        return vId;
    }
}
