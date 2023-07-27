load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "56d8c5a5c91e1af73eca71a6fab2ced959b67c86d12ba37feedb0a2dfea441a6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("//:deps.bzl", "go_dependencies")

go_repository(
    name = "com_github_aclements_go_moremath",
    importpath = "github.com/aclements/go-moremath",
    sum = "h1:xlwdaKcTNVW4PtpQb8aKA4Pjy0CdJHEqvFbAnvR5m2g=",
    version = "v0.0.0-20210112150236-f10218a38794",
)

go_repository(
    name = "com_github_alecthomas_kingpin_v2",
    importpath = "github.com/alecthomas/kingpin/v2",
    sum = "h1:H0aULhgmSzN8xQ3nX1uxtdlTHYoPLu5AhHxWrKI6ocU=",
    version = "v2.3.2",
)

go_repository(
    name = "com_github_andybalholm_brotli",
    importpath = "github.com/andybalholm/brotli",
    sum = "h1:8uQZIdzKmjc/iuPu7O2ioW48L81FgatrcpfFmiq/cCs=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_aymerick_douceur",
    importpath = "github.com/aymerick/douceur",
    sum = "h1:Mv+mAeH1Q+n9Fr+oyamOlAkUNPWPlA8PPGR0QAaYuPk=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_bits_and_blooms_bitset",
    importpath = "github.com/bits-and-blooms/bitset",
    sum = "h1:YjAGVd3XmtK9ktAbX8Zg2g2PwLIMjGREZJHlV4j7NEo=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:9F2/+DoOYIOksmaJFPw1tGFy1eDnIJXg+UHjuD8lTak=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_cloudykit_fastprinter",
    importpath = "github.com/CloudyKit/fastprinter",
    sum = "h1:sR+/8Yb4slttB4vD+b9btVEnWgL3Q00OBTzVT8B9C0c=",
    version = "v0.0.0-20200109182630-33d98a066a53",
)

go_repository(
    name = "com_github_cloudykit_jet_v6",
    importpath = "github.com/CloudyKit/jet/v6",
    sum = "h1:EpcZ6SR9n28BUGtNJSvlBqf90IpjeFr36Tizxhn/oME=",
    version = "v6.2.0",
)

go_repository(
    name = "com_github_cockroachdb_datadriven",
    importpath = "github.com/cockroachdb/datadriven",
    sum = "h1:otljaYPt5hWxV3MUfO5dFPFiOXg9CyG5/kCfayTqsJ4=",
    version = "v1.0.3-0.20230413201302-be42291fc80f",
)

go_repository(
    name = "com_github_cockroachdb_errors",
    importpath = "github.com/cockroachdb/errors",
    sum = "h1:lfxS8zZz1+OjtV4MtNWgboi/W5tyLEB6VQZBXN+0VUU=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_cockroachdb_logtags",
    importpath = "github.com/cockroachdb/logtags",
    sum = "h1:r6VH0faHjZeQy818SGhaone5OnYfxFR/+AzdY3sf5aE=",
    version = "v0.0.0-20230118201751-21c54148d20b",
)

go_repository(
    name = "com_github_cockroachdb_pebble",
    importpath = "github.com/cockroachdb/pebble",
    sum = "h1:lhfCoSWtRiSmod1m9GNs9JSgNk7DGt5B5CxfsBMnAJc=",
    version = "v0.0.0-20230725232348-7a4ed3c57cc5",
)

go_repository(
    name = "com_github_cockroachdb_redact",
    importpath = "github.com/cockroachdb/redact",
    sum = "h1:u1PMllDkdFfPWaNGMyLD1+so+aq3uUItthCFqzwPJ30=",
    version = "v1.1.5",
)

go_repository(
    name = "com_github_cockroachdb_sentry_go",
    importpath = "github.com/cockroachdb/sentry-go",
    sum = "h1:IKgmqgMQlVJIZj19CdocBeSfSaiCbEBZGKODaixqtHM=",
    version = "v0.6.1-cockroachdb.2",
)

go_repository(
    name = "com_github_cockroachdb_tokenbucket",
    importpath = "github.com/cockroachdb/tokenbucket",
    sum = "h1:DJK8W/iB+s/qkTtmXSrHA49lp5O3OsR7E6z4byOLy34=",
    version = "v0.0.0-20230613231145-182959a1fad6",
)

go_repository(
    name = "com_github_codegangsta_inject",
    importpath = "github.com/codegangsta/inject",
    sum = "h1:sDMmm+q/3+BukdIpxwO365v/Rbspp2Nt5XntgQRXq8Q=",
    version = "v0.0.0-20150114235600-33e0aa1cb7c0",
)

go_repository(
    name = "com_github_consensys_bavard",
    importpath = "github.com/consensys/bavard",
    sum = "h1:oLhMLOFGTLdlda/kma4VOJazblc7IM5y5QPd2A/YjhQ=",
    version = "v0.1.13",
)

go_repository(
    name = "com_github_crate_crypto_go_ipa",
    importpath = "github.com/crate-crypto/go-ipa",
    sum = "h1:6IrxszG5G+O7zhtkWxq6+unVvnrm1fqV2Pe+T95DUzw=",
    version = "v0.0.0-20220523130400-f11357ae11c7",
)

go_repository(
    name = "com_github_crate_crypto_go_kzg_4844",
    importpath = "github.com/crate-crypto/go-kzg-4844",
    sum = "h1:UVuHOE+5tIWrim4zf/Xaa43+MIsDCPyW76QhUpiMGj4=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_datadog_zstd",
    importpath = "github.com/DataDog/zstd",
    sum = "h1:oWf5W7GtOLgp6bciQYDmhHHjdhYkALu6S/5Ni9ZgSvQ=",
    version = "v1.5.5",
)

go_repository(
    name = "com_github_deckarep_golang_set_v2",
    importpath = "github.com/deckarep/golang-set/v2",
    sum = "h1:qs18EKUfHm2X9fA50Mr/M5hccg2tNnVqsiBImnyDs0g=",
    version = "v2.3.0",
)

go_repository(
    name = "com_github_eknkc_amber",
    importpath = "github.com/eknkc/amber",
    sum = "h1:clC1lXBpe2kTj2VHdaIu9ajZQe4kcEY9j0NsnDDBZ3o=",
    version = "v0.0.0-20171010120322-cdade1c07385",
)

go_repository(
    name = "com_github_ethereum_c_kzg_4844",
    importpath = "github.com/ethereum/c-kzg-4844",
    sum = "h1:+cUvymlnoDDQgMInp25Bo3OmLajmmY8mLJ/tLjqd77Q=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_fatih_structs",
    importpath = "github.com/fatih/structs",
    sum = "h1:Q7juDM0QtcnhCpeyLGQKyg4TOIghuNXrkL32pHAUMxo=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_flosch_pongo2_v4",
    importpath = "github.com/flosch/pongo2/v4",
    sum = "h1:gv+5Pe3vaSVmiJvh/BZa82b7/00YUGm0PIyVVLop0Hw=",
    version = "v4.0.2",
)

go_repository(
    name = "com_github_gballet_go_verkle",
    importpath = "github.com/gballet/go-verkle",
    sum = "h1:AB7YjNrzlVHsYz06zCULVV2zYCEft82P86dSmtwxKL0=",
    version = "v0.0.0-20220902153445-097bd83b7732",
)

go_repository(
    name = "com_github_getsentry_sentry_go",
    importpath = "github.com/getsentry/sentry-go",
    sum = "h1:XNX9zKbv7baSEI65l+H1GEJgSeIC1c7EN5kluWaP6dM=",
    version = "v0.22.0",
)

go_repository(
    name = "com_github_ghemawat_stream",
    importpath = "github.com/ghemawat/stream",
    sum = "h1:r5GgOLGbza2wVHRzK7aAj6lWZjfbAwiu/RDCVOKjRyM=",
    version = "v0.0.0-20171120220530-696b145b53b9",
)

go_repository(
    name = "com_github_gin_contrib_sse",
    importpath = "github.com/gin-contrib/sse",
    sum = "h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_gin_gonic_gin",
    importpath = "github.com/gin-gonic/gin",
    sum = "h1:4+fr/el88TOO3ewCmQr8cx/CtZ/umlIRIs5M4NTNjf8=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_go_errors_errors",
    importpath = "github.com/go-errors/errors",
    sum = "h1:J6MZopCL4uSllY1OfXM374weqZFFItUbrImctkmUxIA=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_go_kit_log",
    importpath = "github.com/go-kit/log",
    sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_go_martini_martini",
    importpath = "github.com/go-martini/martini",
    sum = "h1:xveKWz2iaueeTaUgdetzel+U7exyigDYBryyVfV/rZk=",
    version = "v0.0.0-20170121215854-22fa46961aab",
)

go_repository(
    name = "com_github_go_playground_locales",
    importpath = "github.com/go-playground/locales",
    sum = "h1:u50s323jtVGugKlcYeyzC0etD1HifMjqmJqb8WugfUU=",
    version = "v0.14.0",
)

go_repository(
    name = "com_github_go_playground_universal_translator",
    importpath = "github.com/go-playground/universal-translator",
    sum = "h1:82dyy6p4OuJq4/CByFNOn/jYrnRPArHwAcmLoJZxyho=",
    version = "v0.18.0",
)

go_repository(
    name = "com_github_go_playground_validator_v10",
    importpath = "github.com/go-playground/validator/v10",
    sum = "h1:prmOlTVv+YjZjmRmNSF3VmspqJIxJWXmqUsHwfTRRkQ=",
    version = "v10.11.1",
)

go_repository(
    name = "com_github_goccy_go_json",
    importpath = "github.com/goccy/go-json",
    sum = "h1:/pAaQDLHEoCq/5FFmSKBswWmK6H0e8g4159Kc/X/nqk=",
    version = "v0.9.11",
)

go_repository(
    name = "com_github_gofrs_flock",
    importpath = "github.com/gofrs/flock",
    sum = "h1:+gYjHKf32LDeiEEFhQaotPbLuUXjY5ZqxKgXy7n59aw=",
    version = "v0.8.1",
)

go_repository(
    name = "com_github_gogo_googleapis",
    importpath = "github.com/gogo/googleapis",
    sum = "h1:1Yx4Myt7BxzvUr5ldGSbwYiZG6t9wGBZ+8/fX3Wvtq0=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_gogo_status",
    importpath = "github.com/gogo/status",
    sum = "h1:+eIkrewn5q6b30y+g/BJINVVdi2xH7je5MPJ3ZPK3JA=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_gorilla_css",
    importpath = "github.com/gorilla/css",
    sum = "h1:BQqNyPTi50JCFMTw/b67hByjMVXZRwGha6wxVGkeihY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_guptarohit_asciigraph",
    importpath = "github.com/guptarohit/asciigraph",
    sum = "h1:ccFnUF8xYIOUPPY3tmdvRyHqmn1MYI9iv1pLKX+/ZkQ=",
    version = "v0.5.5",
)

go_repository(
    name = "com_github_hdrhistogram_hdrhistogram_go",
    importpath = "github.com/HdrHistogram/hdrhistogram-go",
    sum = "h1:5IcZpTvzydCQeHzK4Ef/D5rrSqwxob0t8PQPMybUNFM=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_hydrogen18_memlistener",
    importpath = "github.com/hydrogen18/memlistener",
    sum = "h1:JR7eDj8HD6eXrc5fWLbSUnfcQFL06PYvCc0DKQnWfaU=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_influxdata_influxdb1_client",
    importpath = "github.com/influxdata/influxdb1-client",
    sum = "h1:qSHzRbhzK8RdXOsAdfDgO49TtqC1oZ+acxPrkfTxcCs=",
    version = "v0.0.0-20220302092344-a9ab5670611c",
)

go_repository(
    name = "com_github_iris_contrib_schema",
    importpath = "github.com/iris-contrib/schema",
    sum = "h1:CPSBLyx2e91H2yJzPuhGuifVRnZBBJ3pCOMbOvPZaTw=",
    version = "v0.0.6",
)

go_repository(
    name = "com_github_joker_jade",
    importpath = "github.com/Joker/jade",
    sum = "h1:Qbeh12Vq6BxURXT1qZBRHsDxeURB8ztcL6f3EXSGeHk=",
    version = "v1.1.3",
)

go_repository(
    name = "com_github_josharian_intern",
    importpath = "github.com/josharian/intern",
    sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jpillora_backoff",
    importpath = "github.com/jpillora/backoff",
    sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_kataras_blocks",
    importpath = "github.com/kataras/blocks",
    sum = "h1:cF3RDY/vxnSRezc7vLFlQFTYXG/yAr1o7WImJuZbzC4=",
    version = "v0.0.7",
)

go_repository(
    name = "com_github_kataras_golog",
    importpath = "github.com/kataras/golog",
    sum = "h1:isP8th4PJH2SrbkciKnylaND9xoTtfxv++NB+DF0l9g=",
    version = "v0.1.8",
)

go_repository(
    name = "com_github_kataras_iris_v12",
    importpath = "github.com/kataras/iris/v12",
    sum = "h1:WzDY5nGuW/LgVaFS5BtTkW3crdSKJ/FEgWnxPnIVVLI=",
    version = "v12.2.0",
)

go_repository(
    name = "com_github_kataras_pio",
    importpath = "github.com/kataras/pio",
    sum = "h1:kqreJ5KOEXGMwHAWHDwIl+mjfNCPhAwZPa8gK7MKlyw=",
    version = "v0.0.11",
)

go_repository(
    name = "com_github_kataras_sitemap",
    importpath = "github.com/kataras/sitemap",
    sum = "h1:w71CRMMKYMJh6LR2wTgnk5hSgjVNB9KL60n5e2KHvLY=",
    version = "v0.0.6",
)

go_repository(
    name = "com_github_kataras_tunnel",
    importpath = "github.com/kataras/tunnel",
    sum = "h1:sCAqWuJV7nPzGrlb0os3j49lk2JhILT0rID38NHNLpA=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_kilic_bls12_381",
    importpath = "github.com/kilic/bls12-381",
    sum = "h1:encrdjqKMEvabVQ7qYOKu1OvhqpK4s47wDYtNiPtlp4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:2mk3MPGNzKyxErAw8YaohYh69+pa4sIQSC0fPGCFR9I=",
    version = "v1.16.7",
)

go_repository(
    name = "com_github_labstack_echo_v4",
    importpath = "github.com/labstack/echo/v4",
    sum = "h1:5CiyngihEO4HXsz3vVsJn7f8xAlWwRr3aY6Ih280ZKA=",
    version = "v4.10.0",
)

go_repository(
    name = "com_github_labstack_gommon",
    importpath = "github.com/labstack/gommon",
    sum = "h1:y7cvthEAEbU0yHOf4axH8ZG2NH8knB9iNSoTO8dyIk8=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_leodido_go_urn",
    importpath = "github.com/leodido/go-urn",
    sum = "h1:BqpAaACuzVSgi/VLzGZIobT2z4v53pjosyNd9Yv6n/w=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_mailgun_raymond_v2",
    importpath = "github.com/mailgun/raymond/v2",
    sum = "h1:5dmlB680ZkFG2RN/0lvTAghrSxIESeu9/2aeDqACtjw=",
    version = "v2.0.48",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
    version = "v0.7.7",
)

go_repository(
    name = "com_github_microcosm_cc_bluemonday",
    importpath = "github.com/microcosm-cc/bluemonday",
    sum = "h1:SMZe2IGa0NuHvnVNAZ+6B38gsTbi5e4sViiWJyDDqFY=",
    version = "v1.0.23",
)

go_repository(
    name = "com_github_mmcloughlin_addchain",
    importpath = "github.com/mmcloughlin/addchain",
    sum = "h1:SobOdjm2xLj1KkXN5/n0xTIWyZA2+s99UCY1iPfkHRY=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
    version = "v0.0.0-20190716064945-2f068394615f",
)

go_repository(
    name = "com_github_pelletier_go_toml_v2",
    importpath = "github.com/pelletier/go-toml/v2",
    sum = "h1:ipoSadvV8oGUjnUbMub59IDPPwfxF694nG/jwbMiyQg=",
    version = "v2.0.5",
)

go_repository(
    name = "com_github_pingcap_errors",
    importpath = "github.com/pingcap/errors",
    sum = "h1:lFuQV/oaUMGcD2tqt+01ROSmJs75VG1ToEOkZIZ4nE4=",
    version = "v0.11.4",
)

go_repository(
    name = "com_github_protolambda_bls12_381_util",
    importpath = "github.com/protolambda/bls12-381-util",
    sum = "h1:cZC+usqsYgHtlBaGulVnZ1hfKAi8iWtujBnRLQE698c=",
    version = "v0.0.0-20220416220906-d8552aa452c7",
)

go_repository(
    name = "com_github_schollz_closestmatch",
    importpath = "github.com/schollz/closestmatch",
    sum = "h1:Uel2GXEpJqOWBrlyI+oY9LTiyyjYS17cCYRqP13/SHk=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_shopify_goreferrer",
    importpath = "github.com/Shopify/goreferrer",
    sum = "h1:KkH3I3sJuOLP3TjA/dfr4NAY8bghDwnXiU7cTKxQqo0=",
    version = "v0.0.0-20220729165902-8cddb4f5de06",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:trlNQbNUG3OdDrDil03MCb1H2o9nJ1x4/5LYw7byDE0=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    sum = "h1:6m/oheQuQ13N9ks4hubMG6BnvwOeaJrqSPLahSnczz8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_tdewolff_minify_v2",
    importpath = "github.com/tdewolff/minify/v2",
    sum = "h1:kejsHQMM17n6/gwdw53qsi6lg0TGddZADVyQOz1KMdE=",
    version = "v2.12.4",
)

go_repository(
    name = "com_github_tdewolff_parse_v2",
    importpath = "github.com/tdewolff/parse/v2",
    sum = "h1:KCkDvNUMof10e3QExio9OPZJT8SbdKojLBumw8YZycQ=",
    version = "v2.6.4",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:YPXUKf7fYbp/y8xloBqZOw2qaVggbfwMlI8WM3wZUJ0=",
    version = "v1.2.7",
)

go_repository(
    name = "com_github_urfave_negroni",
    importpath = "github.com/urfave/negroni",
    sum = "h1:kIimOitoypq34K7TG7DUaJ9kq/N4Ofuwi1sjz0KipXc=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    importpath = "github.com/valyala/bytebufferpool",
    sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasthttp",
    importpath = "github.com/valyala/fasthttp",
    sum = "h1:CRq/00MfruPGFLTQKY8b+8SfdK60TxNztjRMnH0t1Yc=",
    version = "v1.40.0",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    importpath = "github.com/valyala/fasttemplate",
    sum = "h1:lxLXG0uE3Qnshl9QyaK6XJxMXlQZELvChBOCmQD0Loo=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_vmihailenco_msgpack_v5",
    importpath = "github.com/vmihailenco/msgpack/v5",
    sum = "h1:5gO0H1iULLWGhs2H5tbAHIZTV8/cYafcFOr9znI5mJU=",
    version = "v5.3.5",
)

go_repository(
    name = "com_github_vmihailenco_tagparser_v2",
    importpath = "github.com/vmihailenco/tagparser/v2",
    sum = "h1:y09buUbR+b5aycVFQs/g70pqKVZNBmxwAhO7/IwNM9g=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_xhit_go_str2duration_v2",
    importpath = "github.com/xhit/go-str2duration/v2",
    sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_yosssi_ace",
    importpath = "github.com/yosssi/ace",
    sum = "h1:tUkIP/BLdKqrlrPwcmH0shwEEhTRHoGnc1wFIWmaBUA=",
    version = "v0.0.5",
)

go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    sum = "h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=",
    version = "v1.67.0",
)

go_repository(
    name = "in_gopkg_natefinch_lumberjack_v2",
    importpath = "gopkg.in/natefinch/lumberjack.v2",
    sum = "h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=",
    version = "v2.0.0",
)

go_repository(
    name = "io_rsc_tmplfunc",
    importpath = "rsc.io/tmplfunc",
    sum = "h1:53XFQh69AfOa8Tw0Jm7t+GV7KZhOi6jzsCzTtKbMvzU=",
    version = "v0.0.3",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
    version = "v1.6.7",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:rtNKfB++wz5mtDY2t5C8TXlU5y52ojSu7tZo0z7u8eQ=",
    version = "v0.0.0-20230227214838-9b19f0bdc514",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:LAv2ds7cmFV/XTS3XG1NneeENYrXGmorPxsBbptIjNc=",
    version = "v1.53.0",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:Di6/M8l0O2lCLc6VVRWhgCiApHV8MnQurBnFSHsQtNY=",
    version = "v0.0.0-20230725093048-515e97ebf090",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:6dkIjl3j3LtZ/O3sTgZTMsLKSftL/B8Zgq4huOIIUu8=",
    version = "v0.8.0",
)

go_repository(
    name = "org_golang_x_perf",
    importpath = "golang.org/x/perf",
    sum = "h1:ObuXPmIgI4ZMyQLIz48cJYgSyWdjUXc2SZAdyJMwEAU=",
    version = "v0.0.0-20230113213139-801c7ef9e5c5",
)

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:nogo",
    version = "1.19.7",
)

gazelle_dependencies()
