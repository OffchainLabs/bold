load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_alecthomas_units",
        importpath = "github.com/alecthomas/units",
        sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
        version = "v0.0.0-20211218093645-b94a6e3cc137",
    )

    go_repository(
        name = "com_github_allegro_bigcache",
        importpath = "github.com/allegro/bigcache",
        sum = "h1:eMwmnE/GDgah4HI848JfFxHt+iPb26b4zyfspmqY0/8=",
        version = "v1.2.1-0.20190218064605-e24eb225f156",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:BS+UYpbsElC82gB+2E2jiCBg36i8HlubTB/dO/moQ9c=",
        version = "v1.2.0",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_config",
        importpath = "github.com/aws/aws-sdk-go-v2/config",
        sum = "h1:ZAoq32boMzcaTW9bcUacBswAmHTbvlvDJICgHFZuECo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_credentials",
        importpath = "github.com/aws/aws-sdk-go-v2/credentials",
        sum = "h1:NbvWIM1Mx6sNPTxowHgS2ewXCRp+NGTzUYb/96FZJbY=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
        sum = "h1:EtEU7WRaWliitZh2nmuxEXrN0Cb8EgPUFGIoTMeqbzI=",
        version = "v1.0.2",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
        sum = "h1:4AH9fFjUlVktQMznF+YN33aWNXaR4VgDXyP28qokJC0=",
        version = "v1.0.2",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_route53",
        importpath = "github.com/aws/aws-sdk-go-v2/service/route53",
        sum = "h1:cKr6St+CtC3/dl/rEBJvlk7A/IN5D5F02GNkGzfbtVU=",
        version = "v1.1.1",
    )

    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sso",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
        sum = "h1:37QubsarExl5ZuCBlnRP+7l1tNwZPBSTqpTBrPH98RU=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sts",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
        sum = "h1:TJoIfnIFubCX0ACVeJ0w46HEH5MwjwYN4iFhuYIhfIY=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:D6CSsM3gdxaGaqXnPgOBCeL6Mophqzu7KJOu7zW78sU=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_azcore",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/azcore",
        sum = "h1:qoVeMsc9/fh/yhxVaA0obYjVH/oI/ihrOoMwsLS9KSA=",
        version = "v0.21.1",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_internal",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/internal",
        sum = "h1:E+m3SkZCN0Bf5q7YdTs5lSm2CYY3CK4spn5OmUIiQtk=",
        version = "v0.8.3",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_storage_azblob",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob",
        sum = "h1:Px2UA+2RvSSvv+RvJNuUB6n7rs5Wsel4dXLe90Um2n4=",
        version = "v0.3.0",
    )

    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )

    go_repository(
        name = "com_github_btcsuite_btcd_btcec_v2",
        importpath = "github.com/btcsuite/btcd/btcec/v2",
        sum = "h1:5n0X6hX0Zk+6omWcihdYvdAlGf2DfasC0GMf7DClJ3U=",
        version = "v2.3.2",
    )
    go_repository(
        name = "com_github_btcsuite_btcd_chaincfg_chainhash",
        importpath = "github.com/btcsuite/btcd/chaincfg/chainhash",
        sum = "h1:q0rUy8C/TYNBQS1+CGKw68tLOFYSNEs0TFnxxnS9+4U=",
        version = "v1.0.1",
    )

    go_repository(
        name = "com_github_cespare_cp",
        importpath = "github.com/cespare/cp",
        sum = "h1:SE+dxFebS7Iik5LK0tsi1k9ZCxEaFX4AjQmoyA+1dJk=",
        version = "v0.1.0",
    )

    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_cloudflare_cloudflare_go",
        importpath = "github.com/cloudflare/cloudflare-go",
        sum = "h1:gFqGlGl/5f9UGXAaKapCGUfaTCgRKKnzu2VvzMZlOFA=",
        version = "v0.14.0",
    )

    go_repository(
        name = "com_github_consensys_gnark_crypto",
        importpath = "github.com/consensys/gnark-crypto",
        sum = "h1:zRh22SR7o4K35SoNqouS9J/TKHTyU2QWaj5ldehyXtA=",
        version = "v0.10.0",
    )

    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        sum = "h1:p1EgwI/C7NhT0JmVkwCD2ZBK8j4aeHQX2pMHHBfMQ6w=",
        version = "v2.0.2",
    )

    go_repository(
        name = "com_github_creack_pty",
        importpath = "github.com/creack/pty",
        sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_d4l3k_messagediff",
        importpath = "github.com/d4l3k/messagediff",
        sum = "h1:ZcAIMYsUg0EAp9X+tt8/enBE/Q8Yd5kzPynLyKptt9U=",
        version = "v1.2.1",
    )

    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )

    go_repository(
        name = "com_github_decred_dcrd_crypto_blake256",
        importpath = "github.com/decred/dcrd/crypto/blake256",
        sum = "h1:7PltbUIQB7u/FfZ39+DGa/ShuMyJ5ilcvdfma9wOH6Y=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_decred_dcrd_dcrec_secp256k1_v4",
        importpath = "github.com/decred/dcrd/dcrec/secp256k1/v4",
        sum = "h1:8UrgZ3GkP4i/CLijOJx79Yu+etlyjdBU4sfcs2WYQMs=",
        version = "v4.2.0",
    )
    go_repository(
        name = "com_github_deepmap_oapi_codegen",
        importpath = "github.com/deepmap/oapi-codegen",
        sum = "h1:SegyeYGcdi0jLLrpbCMoJxnUUn8GBXHsvr4rbzjuhfU=",
        version = "v1.8.2",
    )

    go_repository(
        name = "com_github_dlclark_regexp2",
        importpath = "github.com/dlclark/regexp2",
        sum = "h1:7lJfhqlPssTb1WQx4yvTHN0uElPEv52sbaECrAQxjAo=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        sum = "h1:HlFGsy+9/xrgMmhmN+NGhCc5SHGJ7I+kHosRR1xc/aI=",
        version = "v1.6.2",
    )

    go_repository(
        name = "com_github_dop251_goja",
        importpath = "github.com/dop251/goja",
        sum = "h1:kgvzE5wLsLa7XKfV85VZl40QXaMCaeFtHpPwJ8fhotY=",
        version = "v0.0.0-20230122112309-96b1610dd4f7",
    )

    go_repository(
        name = "com_github_ethereum_go_ethereum",
        importpath = "github.com/ethereum/go-ethereum",
        patch_args = ["-p1"],
        patches = [
            "//third_party:com_github_ethereum_go_ethereum_secp256k1.patch",
        ],
        sum = "h1:bdnhLPtqETd4m3mS8BGMNvBTf36bO5bx/hxE2zljOa0=",
        version = "v1.12.0",
    )

    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_fjl_gencodec",
        importpath = "github.com/fjl/gencodec",
        sum = "h1:bBLctRc7kr01YGvaDfgLbTwjFNW5jdp5y5rj8XXBHfY=",
        version = "v0.0.0-20230517082657-f9840df7b83e",
    )

    go_repository(
        name = "com_github_fjl_memsize",
        importpath = "github.com/fjl/memsize",
        sum = "h1:FtmdgXiUlNeRsoNMFlKLDt+S+6hbjVMEW6RGQ7aUf7c=",
        version = "v0.0.0-20190710130421-bcb5799ab5e5",
    )

    go_repository(
        name = "com_github_fsnotify_fsnotify",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:n+5WquG0fcWoWp6xPWfHdbskMCQaFnG6PfBrh1Ky4HY=",
        version = "v1.6.0",
    )

    go_repository(
        name = "com_github_garslo_gogen",
        importpath = "github.com/garslo/gogen",
        sum = "h1:IZqZOB2fydHte3kUgxrzK5E1fW7RQGeDwE8F/ZZnUYc=",
        version = "v0.0.0-20170306192744-1d203ffc1f61",
    )

    go_repository(
        name = "com_github_gballet_go_libpcsclite",
        importpath = "github.com/gballet/go-libpcsclite",
        sum = "h1:f6D9Hr8xV8uYKlyuj8XIruxlh9WjVjdh1gIicAS7ays=",
        version = "v0.0.0-20191108122812-4678299bea08",
    )

    go_repository(
        name = "com_github_go_logfmt_logfmt",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )

    go_repository(
        name = "com_github_go_ole_go_ole",
        importpath = "github.com/go-ole/go-ole",
        sum = "h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=",
        version = "v1.2.6",
    )

    go_repository(
        name = "com_github_go_sourcemap_sourcemap",
        importpath = "github.com/go-sourcemap/sourcemap",
        sum = "h1:W1iEw64niKVGogNgBN3ePyLFfuisuzeidWPMPWmECqU=",
        version = "v2.1.3+incompatible",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        sum = "h1:ntEHSVwIt7PNXNpgPmVfMrNhLtgjlmnZha2kOpuRiDw=",
        version = "v1.8.1",
    )

    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )

    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:kHL1vqdqWNfATmA0FNMdmZNMyZI1U6O31X4rlIPoBog=",
        version = "v4.3.0",
    )

    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        sum = "h1:PBC98N2aIaM3XXiurYmW7fx4GZkL8feAMVq7nEjURHk=",
        version = "v0.0.5-0.20220116011046-fa5810519dcb",
    )

    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
        version = "v0.5.9",
    )

    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        sum = "h1:Q75Upo5UN4JbPFURXZ8nLKYUvF85dyFRop/vQ0Rv+64=",
        version = "v1.1.1-0.20200604201612-c04b05f3adfa",
    )

    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sum = "h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=",
        version = "v1.8.0",
    )

    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_graph_gophers_graphql_go",
        importpath = "github.com/graph-gophers/graphql-go",
        sum = "h1:Eb9x/q6MFpCLz7jBCiP/WTxjSDrYLR1QY41SORZyNJ0=",
        version = "v1.3.0",
    )

    go_repository(
        name = "com_github_hashicorp_go_bexpr",
        importpath = "github.com/hashicorp/go-bexpr",
        sum = "h1:9kuI5PFotCboP3dkDYFr/wi0gg0QVbSNz5oFRpxn4uE=",
        version = "v0.1.10",
    )

    go_repository(
        name = "com_github_holiman_bloomfilter_v2",
        importpath = "github.com/holiman/bloomfilter/v2",
        sum = "h1:73e0e/V0tCydx14a0SCYS/EWCxgwLZ18CZcZKVu0fao=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_holiman_uint256",
        importpath = "github.com/holiman/uint256",
        sum = "h1:K8UWO1HUJpRMXBxbmaY1Y8IAMZC/RsKB+ArEnnK4l5o=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        importpath = "github.com/hpcloud/tail",
        sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_huin_goupnp",
        importpath = "github.com/huin/goupnp",
        sum = "h1:uOKW26NG1hsSSbXIZ1IR7XP9Gjd1U8pnLaCMgntmkmY=",
        version = "v1.2.0",
    )

    go_repository(
        name = "com_github_influxdata_influxdb_client_go_v2",
        importpath = "github.com/influxdata/influxdb-client-go/v2",
        sum = "h1:HGBfZYStlx3Kqvsv1h2pJixbCl/jhnFtxpKFAv9Tu5k=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_influxdata_line_protocol",
        importpath = "github.com/influxdata/line-protocol",
        sum = "h1:vilfsDSy7TDxedi9gyBkMvAirat/oRcL0lFdJBf6tdM=",
        version = "v0.0.0-20210311194329-9aa0e372d097",
    )

    go_repository(
        name = "com_github_jackpal_go_nat_pmp",
        importpath = "github.com/jackpal/go-nat-pmp",
        sum = "h1:KzKSgb7qkJvOUTqYl9/Hg/me3pWgBmERKrTGD7BdWus=",
        version = "v1.0.2",
    )

    go_repository(
        name = "com_github_jedisct1_go_minisign",
        importpath = "github.com/jedisct1/go-minisign",
        sum = "h1:UvSe12bq+Uj2hWd8aOlwPmoZ+CITRFrdit+sDGfAg8U=",
        version = "v0.0.0-20190909160543-45766022959e",
    )

    go_repository(
        name = "com_github_julienschmidt_httprouter",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_karalabe_usb",
        importpath = "github.com/karalabe/usb",
        sum = "h1:M6QQBNxF+CQ8OFvxrT90BA0qBOXymndZnk5q235mFc4=",
        version = "v0.0.2",
    )

    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )

    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
        version = "v1.1.0",
    )

    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:BTarxUcIeDqL27Mc+vyvdWYSL28zpIhv3RoTdsLMPng=",
        version = "v0.0.17",
    )

    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:UNAjwbU9l54TA3KzvqLGxwWjHmMgBUVhBiTjelZgg3U=",
        version = "v0.0.15",
    )

    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
        version = "v1.0.4",
    )

    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:CpVNEelQCZBooIPDn+AR3NpivK/TIKU8bDxdASFVQag=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_mitchellh_pointerstructure",
        importpath = "github.com/mitchellh/pointerstructure",
        sum = "h1:O+i9nHnXS3l/9Wu7r4NrEdwA2VFTicjUEN1uBnDo34A=",
        version = "v1.2.0",
    )

    go_repository(
        name = "com_github_naoina_go_stringutil",
        importpath = "github.com/naoina/go-stringutil",
        sum = "h1:rCUeRUHjBjGTSHl0VC00jUPLz8/F9dDzYI70Hzifhks=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_naoina_toml",
        importpath = "github.com/naoina/toml",
        sum = "h1:shk/vn9oCoOTmwcouEdwIeOtOGA/ELRUw/GwvxwfT+0=",
        version = "v0.1.2-0.20170918210437-9fafd6967416",
    )

    go_repository(
        name = "com_github_nxadm_tail",
        importpath = "github.com/nxadm/tail",
        sum = "h1:DQuhQpB1tVlglWS2hLQ5OV6B5r8aGxSrPc5Qo6uTN78=",
        version = "v1.4.4",
    )

    go_repository(
        name = "com_github_olekukonko_tablewriter",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
        version = "v0.0.5",
    )

    go_repository(
        name = "com_github_onsi_ginkgo",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:2mOpI4JVVPBN+WQRa0WKH2eXR+Ey+uK4n7Zj0aYpIQA=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        sum = "h1:o0+MgICZLuZ7xjH7Vx6zS/zcu93/BEp1VwkIW1mEXCE=",
        version = "v1.10.1",
    )

    go_repository(
        name = "com_github_opentracing_opentracing_go",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:pWlfV3Bxv7k65HYwkikxat0+s3pV4bsqf19k25Ur8rU=",
        version = "v1.1.0",
    )

    go_repository(
        name = "com_github_peterh_liner",
        importpath = "github.com/peterh/liner",
        sum = "h1:oYW+YCJ1pachXTQmzR3rNLYGGz4g/UgFcjb28p/viDM=",
        version = "v1.1.1-0.20190123174540-a2c9a5303de7",
    )

    go_repository(
        name = "com_github_pkg_diff",
        importpath = "github.com/pkg/diff",
        sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
        version = "v0.0.0-20210226163009-20ebb0f2a09e",
    )

    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:yk/hx9hDbrGHovbci4BY+pRMfSuuat626eFsHb7tmT8=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:5lQXD3cAg1OXBf4Wq03gTrXHeaV0TQvGfUooCfx1yqY=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sum = "h1:+5BrQJwiBB9xsMygAB3TNvpQKOwlkc25LbISbrdOOfY=",
        version = "v0.44.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:xRC8Iq1yyca5ypa9n1EZnWZkt7dwcoRPQwX/5gwaUuI=",
        version = "v0.11.1",
    )

    go_repository(
        name = "com_github_rivo_uniseg",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:8TfxU8dW6PdqD27gjM8MVNuicgxIjxpm4K7x4jp8sis=",
        version = "v0.4.4",
    )

    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:cWPaGQEPrBb5/AsnsZesgZZ9yb1OQ+GOISoDNXVBh4M=",
        version = "v1.11.0",
    )

    go_repository(
        name = "com_github_rs_cors",
        importpath = "github.com/rs/cors",
        sum = "h1:+88SsELBHx5r+hZ8TCkggzSstaWNbDvThkVK8H6f9ik=",
        version = "v1.7.0",
    )

    go_repository(
        name = "com_github_russross_blackfriday_v2",
        importpath = "github.com/russross/blackfriday/v2",
        sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
        version = "v2.1.0",
    )

    go_repository(
        name = "com_github_shirou_gopsutil",
        importpath = "github.com/shirou/gopsutil",
        sum = "h1:+1+c1VGhc88SSonWP6foOcLhvnKlUeu/erjjvaPEYiI=",
        version = "v3.21.11+incompatible",
    )

    go_repository(
        name = "com_github_stackexchange_wmi",
        importpath = "github.com/StackExchange/wmi",
        sum = "h1:fLjPD/aNc3UIOA6tDi6QXUemppXK3P9BI7mr2hd6gx8=",
        version = "v0.0.0-20180116203802-5d049714c4a6",
    )
    go_repository(
        name = "com_github_status_im_keycard_go",
        importpath = "github.com/status-im/keycard-go",
        sum = "h1:QDLFswOQu1r5jsycloeQh3bVU8n/NatHHaZobtDnDzA=",
        version = "v0.2.0",
    )

    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_github_supranational_blst",
        importpath = "github.com/supranational/blst",
        sum = "h1:u49mjRnygnB34h8OKbnNJFVUtWSKIKb1KukdV8bILUM=",
        version = "v0.3.11-0.20230406105308-e9dfc5ee724b",
    )

    go_repository(
        name = "com_github_syndtr_goleveldb",
        importpath = "github.com/syndtr/goleveldb",
        sum = "h1:epCh84lMvA70Z7CTTCmYQn2CKbY8j86K7/FAIr141uY=",
        version = "v1.0.1-0.20210819022825-2ae1ddf74ef7",
    )

    go_repository(
        name = "com_github_tklauser_go_sysconf",
        importpath = "github.com/tklauser/go-sysconf",
        sum = "h1:89WgdJhk5SNwJfu+GKyYveZ4IaJ7xAkecBo+KdJV0CM=",
        version = "v0.3.11",
    )
    go_repository(
        name = "com_github_tklauser_numcpus",
        importpath = "github.com/tklauser/numcpus",
        sum = "h1:ng9scYS7az0Bk4OZLvrNXNSAO2Pxr1XXRAPyjhIx+Fk=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_tyler_smith_go_bip39",
        importpath = "github.com/tyler-smith/go-bip39",
        sum = "h1:5eUemwrMargf3BSLRRCalXT93Ns6pQJIjYQN2nyfOP8=",
        version = "v1.1.0",
    )

    go_repository(
        name = "com_github_urfave_cli_v2",
        importpath = "github.com/urfave/cli/v2",
        sum = "h1:5SqCsI/2Qya2bCzK15ozrqo2sZxkh0FHynJZOTVoV6Q=",
        version = "v2.17.2-0.20221006022127-8f469abc00aa",
    )

    go_repository(
        name = "com_github_victoriametrics_fastcache",
        importpath = "github.com/VictoriaMetrics/fastcache",
        sum = "h1:i0mICQuojGDL3KblA7wUNlY5lOK6a4bwt3uRKnkZU40=",
        version = "v1.12.1",
    )

    go_repository(
        name = "com_github_xrash_smetrics",
        importpath = "github.com/xrash/smetrics",
        sum = "h1:bAn7/zixMGCfxrRTfdpNzjtPYqr8smhKouy9mxVdGPU=",
        version = "v0.0.0-20201216005158-039620a65673",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )

    go_repository(
        name = "com_github_yusufpapurcu_wmi",
        importpath = "github.com/yusufpapurcu/wmi",
        sum = "h1:E1ctvB7uKFMOJw3fdOW32DwGE9I7t++CRUEMKvFoFiw=",
        version = "v1.2.3",
    )

    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_d4l3k_messagediff_v1",
        importpath = "gopkg.in/d4l3k/messagediff.v1",
        sum = "h1:70AthpjunwzUiarMHyED52mj9UwtAnE89l1Gmrt3EU0=",
        version = "v1.2.1",
    )

    go_repository(
        name = "in_gopkg_fsnotify_v1",
        importpath = "gopkg.in/fsnotify.v1",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version = "v1.4.7",
    )

    go_repository(
        name = "in_gopkg_natefinch_npipe_v2",
        importpath = "gopkg.in/natefinch/npipe.v2",
        sum = "h1:+JknDZhAj8YMt7GC73Ei8pv4MzjDUNPHgQWJdtMAaDU=",
        version = "v2.0.0-20160621034901-c1b8fa8bdcce",
    )

    go_repository(
        name = "in_gopkg_tomb_v1",
        importpath = "gopkg.in/tomb.v1",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version = "v1.0.0-20141024135613-dd632973f1e7",
    )

    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )

    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:g0LDEJHgrBl9N9r17Ru3sqWhkIx2NB67okBHPwC7hs8=",
        version = "v1.31.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:6Ewdq3tDic1mg5xRO4milcWCfMVQhI4NkqWWvqejpuA=",
        version = "v0.11.0",
    )

    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:bUO06HqtnRcc/7l71XBe4WcqTZ+3AH1J59zWDDwLKgU=",
        version = "v0.11.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:X2//UzNDwYmtCLn7To6G58Wr6f5ahEAQgKNzv9Y951M=",
        version = "v0.10.0",
    )

    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:ftCYgMx6zT/asHUrPw8BLLscYtGznsLAnjq5RH9P66E=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:SqMFp9UcQJZa+pmYuAKjd9xq1f0j5rLcDIk0mj4qAsA=",
        version = "v0.10.0",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:3R7pNqamzBraeqj/Tj8qt1aQ2HpmlC+Cx/qL/7hn4/c=",
        version = "v0.10.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:LAntKIrcmeSKERyiOh0XMV39LXS8IE9UL2yP7+f5ij4=",
        version = "v0.11.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:rg5rLMjNzMS1RkNLzCG38eapWhnYLFYXDXj2gOlr8j4=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        replace = "golang.org/x/tools",
        sum = "h1:VveCTK38A2rkS8ZqFY25HIDFscX5X9OoEhJd3quQmXU=",
        version = "v0.1.12",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:5Pf6pFKu98ODmgnpvkJ3kFUOQGGLIzLIkbzUHp47618=",
        version = "v0.0.0-20220517211312-f3a8303e98df",
    )
