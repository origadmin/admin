module origadmin/application/admin

go 1.25.5

replace github.com/armon/go-metrics v0.5.4 => github.com/origadmin/go-metrics v0.5.4

replace github.com/origadmin/runtime v0.2.15 => ../../runtime

replace github.com/origadmin/contrib v1.3.0 => ../../contrib

require (
	entgo.io/ent v0.14.5
	github.com/casbin/casbin/v2 v2.135.0
	github.com/envoyproxy/protoc-gen-validate v1.3.0
	github.com/go-kratos/kratos/v2 v2.9.2
	github.com/goexts/generic v0.14.0
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/google/gnostic v0.7.1
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.7.0
	github.com/mattn/go-sqlite3 v1.14.33 // indirect
	github.com/mojocn/base64Captcha v1.3.8
	github.com/origadmin/contrib v1.3.0
	github.com/origadmin/runtime v0.2.15
	github.com/origadmin/slog-kratos v1.0.5 // indirect
	github.com/origadmin/toolkits/codec v1.3.1
	github.com/origadmin/toolkits/crypto v1.3.0
	github.com/origadmin/toolkits/errors v1.3.1
	github.com/sqlite3ent/sqlite3 v1.40.0
	golang.org/x/net v0.49.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260114163908-3f89685c29c3
	google.golang.org/grpc v1.78.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/air-verse/air v1.64.0
	github.com/bufbuild/buf v1.64.0
	github.com/casbin/redis-watcher/v2 v2.8.0
	github.com/go-kratos/kratos/cmd/kratos/v2 v2.0.0-20260105075216-c7a58ff59f80
	github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2 v2.0.0-20260105075216-c7a58ff59f80
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.4
	github.com/joho/godotenv v1.5.1
	github.com/origadmin/toolkits/i18n v1.3.0
	github.com/stretchr/testify v1.11.1
	golang.org/x/exp v0.0.0-20260112195511-716be5621a96
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.6.0
)

require (
	ariga.io/atlas v1.0.0 // indirect
	buf.build/gen/go/bufbuild/bufplugin/protocolbuffers/go v1.36.11-20250718181942-e35f9b667443.1 // indirect
	buf.build/gen/go/bufbuild/protodescriptor/protocolbuffers/go v1.36.11-20250109164928-1da0de137947.1 // indirect
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.11-20251209175733-2a1774d88802.1 // indirect
	buf.build/gen/go/bufbuild/registry/connectrpc/go v1.19.1-20251202164234-62b14f0b533c.2 // indirect
	buf.build/gen/go/bufbuild/registry/protocolbuffers/go v1.36.11-20251202164234-62b14f0b533c.1 // indirect
	buf.build/gen/go/pluginrpc/pluginrpc/protocolbuffers/go v1.36.11-20241007202033-cf42259fcbfc.1 // indirect
	buf.build/go/app v0.2.0 // indirect
	buf.build/go/bufplugin v0.9.0 // indirect
	buf.build/go/bufprivateusage v0.1.0 // indirect
	buf.build/go/interrupt v1.1.0 // indirect
	buf.build/go/protovalidate v1.1.0 // indirect
	buf.build/go/protoyaml v0.6.0 // indirect
	buf.build/go/spdx v0.2.0 // indirect
	buf.build/go/standard v0.1.0 // indirect
	cel.dev/expr v0.25.1 // indirect
	connectrpc.com/connect v1.19.1 // indirect
	connectrpc.com/otelconnect v0.9.0 // indirect
	dario.cat/mergo v1.0.2 // indirect
	github.com/AlecAivazis/survey/v2 v2.3.7 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20250102033503-faa5f7b0171c // indirect
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ThreeDotsLabs/watermill v1.5.1 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/bep/godartsass/v2 v2.5.0 // indirect
	github.com/bep/golibsass v1.2.0 // indirect
	github.com/bits-and-blooms/bitset v1.24.4 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/bmatcuk/doublestar/v4 v4.9.1 // indirect
	github.com/bufbuild/protocompile v0.14.2-0.20260114160500-16922e24f2b6 // indirect
	github.com/bufbuild/protoplugin v0.0.0-20250218205857-750e09ce93e1 // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/casbin/casbin/v3 v3.9.0 // indirect
	github.com/casbin/govaluate v1.10.0 // indirect
	github.com/catppuccin/go v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/charmbracelet/bubbles v0.21.1-0.20250623103423-23b8fd6302d7 // indirect
	github.com/charmbracelet/bubbletea v1.3.10 // indirect
	github.com/charmbracelet/colorprofile v0.4.1 // indirect
	github.com/charmbracelet/huh v0.8.0 // indirect
	github.com/charmbracelet/lipgloss v1.1.0 // indirect
	github.com/charmbracelet/x/ansi v0.11.3 // indirect
	github.com/charmbracelet/x/cellbuf v0.0.14 // indirect
	github.com/charmbracelet/x/exp/strings v0.0.0-20251215102626-e0db08df7383 // indirect
	github.com/charmbracelet/x/term v0.2.2 // indirect
	github.com/cli/browser v1.3.0 // indirect
	github.com/clipperhouse/displaywidth v0.6.2 // indirect
	github.com/clipperhouse/stringish v0.1.1 // indirect
	github.com/clipperhouse/uax29/v2 v2.3.0 // indirect
	github.com/containerd/errdefs v1.0.0 // indirect
	github.com/containerd/errdefs/pkg v0.3.0 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.18.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dchest/uniuri v1.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/cli v29.1.5+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker v28.5.2+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.9.5 // indirect
	github.com/docker/go-connections v0.6.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/proto v1.14.2 // indirect
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-chi/chi/v5 v5.2.4 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2 v2.0.0-20260105075216-c7a58ff59f80 // indirect
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20260105075216-c7a58ff59f80 // indirect
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20260105075216-c7a58ff59f80 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/inflect v0.21.2 // indirect
	github.com/go-playground/form/v4 v4.3.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gofrs/flock v0.13.0 // indirect
	github.com/gohugoio/hashstructure v0.6.0 // indirect
	github.com/gohugoio/hugo v0.154.2 // indirect
	github.com/golang-cz/devslog v0.0.15 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/google/cel-go v0.26.1 // indirect
	github.com/google/gnostic-models v0.7.1 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-containerregistry v0.20.7 // indirect
	github.com/google/subcommands v1.2.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/hashicorp/consul/api v1.33.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.23.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jdx/go-netrc v1.0.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/compress v1.18.3 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/lmittmann/tint v1.1.2 // indirect
	github.com/lucasb-eyer/go-colorful v1.3.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20251013123823-9fd1530e3ec3 // indirect
	github.com/lyft/protoc-gen-star/v2 v2.0.4 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.19 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/term v0.5.2 // indirect
	github.com/morikuni/aec v1.1.0 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/termenv v0.16.0 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/olekukonko/cat v0.0.0-20250911104152-50322a0618f6 // indirect
	github.com/olekukonko/errors v1.1.0 // indirect
	github.com/olekukonko/ll v0.1.3 // indirect
	github.com/olekukonko/tablewriter v1.1.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.1 // indirect
	github.com/origadmin/casbin-watcher/v3 v3.0.0-20260121202419-c22241105874 // indirect
	github.com/origadmin/toolkits/identifier/snowflake v1.3.1 // indirect
	github.com/origadmin/toolkits/slogx v1.3.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/petermattis/goid v0.0.0-20260113132338-7c7de50cc741 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.59.0 // indirect
	github.com/redis/go-redis/v9 v9.0.3 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/segmentio/asm v1.2.1 // indirect
	github.com/segmentio/encoding v0.5.3 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/shoenig/go-m1cpu v0.1.7 // indirect
	github.com/sirupsen/logrus v1.9.4 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/stoewer/go-strcase v1.3.1 // indirect
	github.com/tdewolff/parse/v2 v2.8.5 // indirect
	github.com/tetratelabs/wazero v1.11.0 // indirect
	github.com/tidwall/btree v1.8.1 // indirect
	github.com/tklauser/go-sysconf v0.3.16 // indirect
	github.com/tklauser/numcpus v0.11.0 // indirect
	github.com/vbatts/tar-split v0.12.2 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	github.com/zclconf/go-cty v1.16.2 // indirect
	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
	go.lsp.dev/jsonrpc2 v0.10.0 // indirect
	go.lsp.dev/pkg v0.0.0-20210717090340-384b27a52fb2 // indirect
	go.lsp.dev/protocol v0.12.0 // indirect
	go.lsp.dev/uri v0.3.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.64.0 // indirect
	go.opentelemetry.io/otel v1.39.0 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/otel/trace v1.39.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.47.0 // indirect
	golang.org/x/image v0.34.0 // indirect
	golang.org/x/mod v0.32.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/term v0.39.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	golang.org/x/tools v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260114163908-3f89685c29c3 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	modernc.org/libc v1.67.4 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
	modernc.org/sqlite v1.42.2 // indirect
	pluginrpc.com/pluginrpc v0.5.0 // indirect
)
