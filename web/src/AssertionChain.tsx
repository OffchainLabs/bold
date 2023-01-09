import React, { useEffect, useState, useRef } from 'react';
import {graphviz} from 'd3-graphviz';
import { Graphviz } from 'graphviz-react';

interface Data {
  typ: string;
  contents: string;
  vis: string;
  visChallenge: string;
}

export const AssertionsChain = () => {
  const [dot, setDot] = useState<string>('digraph {}');
  const [dotChallenge, setDotChallenge] = useState<string>('digraph {}');

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8000/api/ws');
    socket.onmessage = (event: any) => {
      const item = JSON.parse(event.data) as Data;
      console.log(item);
      setDot(item.vis);
      setDotChallenge(item.visChallenge)
    };
    return () => socket.close();
  }, []);

  return (
    <section className="p-4 break-all">
        {/* <Graphviz dot={dot} options={{'zoom': true, 'width': 1200 }}></Graphviz> */}
        <Graphviz dot={dotChallenge} options={{'zoom': true, 'width': 1200, 'height': 1200 }}></Graphviz>
    </section>
    
  )
};
