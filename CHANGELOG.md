# Changelog

## [0.13.0](https://github.com/rswanson/lfk/compare/v0.13.2...v0.13.0) (2026-06-01)


### ⚠ BREAKING CHANGES

* add multi-strategy right-sizing advisor overlay ([#148](https://github.com/rswanson/lfk/issues/148))
* CrashLoopBackOff investigator overlay
* The monitoring config fallback key has been renamed from "default" to "_global". Users with "default" in their config must update it to "_global".

### Features

* **actions:** add "Go to Node" to the Pod action menu ([#264](https://github.com/rswanson/lfk/issues/264)) ([#269](https://github.com/rswanson/lfk/issues/269)) ([c9190d1](https://github.com/rswanson/lfk/commit/c9190d1dbc87e5e4f0bb956e058f5be02cdd1b9b))
* add bulk action progress counter ([36d4ed3](https://github.com/rswanson/lfk/commit/36d4ed37050440f4a796f60a3e254d813fa22688))
* add configurable data directories (LFK_*_DIR overrides) ([#246](https://github.com/rswanson/lfk/issues/246)) ([4de5317](https://github.com/rswanson/lfk/commit/4de531735ea2d480165f0d19c0760061ce5ec798))
* add context, namespace, kubeconfig flags for [#16](https://github.com/rswanson/lfk/issues/16) ([97b80c2](https://github.com/rswanson/lfk/commit/97b80c2c0813c0e8c7186b2c198135208049446e))
* add flake.nix ([c90effc](https://github.com/rswanson/lfk/commit/c90effcf6503d09f62a85741d2aadce4aeddd2c5))
* add jump-back navigation history ([#249](https://github.com/rswanson/lfk/issues/249)) ([#256](https://github.com/rswanson/lfk/issues/256)) ([c7eae4b](https://github.com/rswanson/lfk/commit/c7eae4ba6828bcb1abca2a50b63fac6aae449835))
* add min_contrast_ratio theme mutator [#39](https://github.com/rswanson/lfk/issues/39) ([9f3eea7](https://github.com/rswanson/lfk/commit/9f3eea7361ff4ec81337fe07d9a23234943265ff))
* add multi-cluster union view with --union-context and --union-set ([#172](https://github.com/rswanson/lfk/issues/172)) ([ba0f405](https://github.com/rswanson/lfk/commit/ba0f4059db6cd6588a4e4d8c9c4ba77fadd543db))
* add multi-strategy right-sizing advisor overlay ([#148](https://github.com/rswanson/lfk/issues/148)) ([5392610](https://github.com/rswanson/lfk/commit/539261090646f1dc94c19dcd3c1b57eca1e7b1bb))
* add negative namespace selection ([#287](https://github.com/rswanson/lfk/issues/287)) ([653d1cd](https://github.com/rswanson/lfk/commit/653d1cd9d83708ffa727580af4cd03fc5343c302))
* add Next column to CronJob preview ([455ab3c](https://github.com/rswanson/lfk/commit/455ab3c9dd6d141e88a008db1246b7ffa04b69e2))
* add Next column to CronJob preview ([a77e968](https://github.com/rswanson/lfk/commit/a77e968041b962bceae6fbd51f88280aaffd9828)), closes [#50](https://github.com/rswanson/lfk/issues/50)
* add options to disable mouse support [#18](https://github.com/rswanson/lfk/issues/18) ([45f2239](https://github.com/rswanson/lfk/commit/45f22398a385b2932b1dde422ac229eec46176c6))
* add secret_lazy_loading config to speed up Secret list ([68a1a0e](https://github.com/rswanson/lfk/commit/68a1a0e2cb2a784383ccc69d97acf18712780b1f))
* add Tail Logs action on x-&gt;l, remap Logs to x-&gt;L [#38](https://github.com/rswanson/lfk/issues/38) ([6fc0724](https://github.com/rswanson/lfk/commit/6fc0724010aaa9fd5394c5adf5a0367bb54fb5b4))
* add UDPRoutes, ReferenceGrants, BackendTLSPolicies to networking sidebar ([3122021](https://github.com/rswanson/lfk/commit/312202174178cd81fcd5ccb2013d87bff82129c2))
* add watch interval option [#19](https://github.com/rswanson/lfk/issues/19) ([d293cc9](https://github.com/rswanson/lfk/commit/d293cc946b2c49cd7a595283bddcf12cf3482053))
* **app:** add read-only mode with per-context [RO] markers ([1b1d9c1](https://github.com/rswanson/lfk/commit/1b1d9c1738db93ea1b82f9979e8eaef51764832a))
* **app:** add read-only mode with per-context [RO] markers ([c148097](https://github.com/rswanson/lfk/commit/c148097833a0771762632562b5b9066c696f7f6d))
* **app:** apply y/Y to multi-selection ([ce71b97](https://github.com/rswanson/lfk/commit/ce71b97e6acba73d8ffc4cd35a1f669212735925))
* **app:** route :export through the Y bulk dispatcher ([d550328](https://github.com/rswanson/lfk/commit/d55032814bfed5d718126e1624bc59ee6929487d))
* **app:** slow watch tick when foreground-idle for 120s ([#5](https://github.com/rswanson/lfk/issues/5)) ([ff39527](https://github.com/rswanson/lfk/commit/ff395278054770287292b335872a9eb97515bb23))
* **app:** tackle PTY pain points from [#81](https://github.com/rswanson/lfk/issues/81) — selection, mux mode, scrollback ([32be754](https://github.com/rswanson/lfk/commit/32be7546df5bf3df871e7a9d9d38a45e912452b1))
* **app:** throttle watch tick when terminal is unfocused ([#4](https://github.com/rswanson/lfk/issues/4)) ([918fdcf](https://github.com/rswanson/lfk/commit/918fdcfd73e9cde91d16bb9baaced698acd83c99))
* **argocd:** add Sync Wave Timeline overlay ([#160](https://github.com/rswanson/lfk/issues/160)) ([3784fc6](https://github.com/rswanson/lfk/commit/3784fc6e3de25fa3774457f487b0c5840e01131f))
* better spacing for browser view ([6654a8b](https://github.com/rswanson/lfk/commit/6654a8b06befe0e8241a878acc13d77a97cbd0d9))
* better support for NO_COLOR mode [#20](https://github.com/rswanson/lfk/issues/20) [#21](https://github.com/rswanson/lfk/issues/21) ([210df11](https://github.com/rswanson/lfk/commit/210df11085fd6e51c9fac106e82a0949bdbfb1ed))
* **clipboard:** support Windows and Wayland via atotto/clipboard ([#195](https://github.com/rswanson/lfk/issues/195)) ([c1871de](https://github.com/rswanson/lfk/commit/c1871de47df2a9597f13c7421980095b5d2d8b2c))
* **clusters:** add per-cluster color coding with title-bar tint ([#124](https://github.com/rswanson/lfk/issues/124)) ([65da3ac](https://github.com/rswanson/lfk/commit/65da3ac010f4b84b4270dc4ce8662243a7171497))
* color palette update notifications (CSI 996/2031) ([#26](https://github.com/rswanson/lfk/issues/26)) ([50f8aee](https://github.com/rswanson/lfk/commit/50f8aeec0f777e5f149885e9042f63833f367b60))
* column toggle overlay applies edits live; Esc discards [#44](https://github.com/rswanson/lfk/issues/44) ([a830b85](https://github.com/rswanson/lfk/commit/a830b8552855565500ba18463b4d9c11fbbc997f))
* **columns + views:** REV, kubectl-parity audit, k9s-style views config ([#271](https://github.com/rswanson/lfk/issues/271)) ([837961c](https://github.com/rswanson/lfk/commit/837961ce44c1de45eca45e17be0a9f916e86dbe2))
* **config:** make kubeconfig discovery directory configurable ([#243](https://github.com/rswanson/lfk/issues/243)) ([71fddf5](https://github.com/rswanson/lfk/commit/71fddf5c43e18019c7affd934a72372403c36210))
* **copy:** open copy-as picker on Y with YAML / JSON / Table options ([#237](https://github.com/rswanson/lfk/issues/237)) ([9f0851d](https://github.com/rswanson/lfk/commit/9f0851d020c96ae40bf50b63b48366ca923fc1b8))
* CrashLoopBackOff investigator overlay ([93d310e](https://github.com/rswanson/lfk/commit/93d310e8eb2b6547c8967749aec7a5e5a318f9ef))
* derive CRD display names from Kind to preserve camel case ([#306](https://github.com/rswanson/lfk/issues/306)) ([c334c85](https://github.com/rswanson/lfk/commit/c334c85b571e9d867aaf223bfbc8f27b0ff30cba)), closes [#301](https://github.com/rswanson/lfk/issues/301)
* **editors:** revamp edit pane — bordered fields + non-shifting cursor ([55b322a](https://github.com/rswanson/lfk/commit/55b322a33113fbff87ec296a51100e8b5fb41a4e))
* **editors:** wire `s` multi-select + Shift+Y format-copy on ConfigMap + Label editors ([44429de](https://github.com/rswanson/lfk/commit/44429def06b4106787b14dcb5fc15d907c8741be))
* **editors:** wire `s` multi-select + Shift+Y format-copy on Secret editor ([f240d35](https://github.com/rswanson/lfk/commit/f240d358e0417e787eaac477b9b3162d42a4b71d))
* enable log preview by default ([9763eb4](https://github.com/rswanson/lfk/commit/9763eb4f3b37e5367fd2510dd2476b717d1f2cd9))
* **endpoints:** surface addresses, ports, and ready/not-ready counts ([d82435c](https://github.com/rswanson/lfk/commit/d82435c651e5b33240dd8b6d8b1b09fb11d07b73))
* energy/perf measurement harness (Phase 1) ([60c3d41](https://github.com/rswanson/lfk/commit/60c3d4132ea600cae16bfc2b62e5d25e4ec55f09))
* **events:** make the events overlay readable by default ([#263](https://github.com/rswanson/lfk/issues/263)) ([#270](https://github.com/rswanson/lfk/issues/270)) ([8bce4ea](https://github.com/rswanson/lfk/commit/8bce4ea2afc704d655e48f9fa7908feaa124bc58))
* extend PgUp/PgDown/Home/End and gg/G across all navigation contexts [#35](https://github.com/rswanson/lfk/issues/35) ([48312a6](https://github.com/rswanson/lfk/commit/48312a6dbad833709eaf6a59be2e99aa6c43b94b))
* **filters:** add Not Running / Not Bound presets and config invert flag ([#230](https://github.com/rswanson/lfk/issues/230)) ([3af1652](https://github.com/rswanson/lfk/commit/3af16528ec28973e4e0893fb47754aa48a050396))
* flake versioning, fail-fast CI, and pre-push tag guard ([c7b522e](https://github.com/rswanson/lfk/commit/c7b522eacb8c0585ca56deb65c9d9a0f32ae34db))
* fuzzy match for command-bar value completions ([4be0f12](https://github.com/rswanson/lfk/commit/4be0f12cfb7e0adb297e3238143ae10f53e813f9)), closes [#27](https://github.com/rswanson/lfk/issues/27)
* graceful shutdown notice with 10s force-quit timeout ([#314](https://github.com/rswanson/lfk/issues/314)) ([68e03f6](https://github.com/rswanson/lfk/commit/68e03f62ce7f8a279112a4f271efbf52228f53f9))
* help screen splits / (search-with-highlight) from f (filter) ([07dc544](https://github.com/rswanson/lfk/commit/07dc544029dfce3a23e3e61572b89cc3b6dbaca2))
* **jobs,cronjobs:** reorder columns, add Suspend column to Jobs ([ef52d25](https://github.com/rswanson/lfk/commit/ef52d255d5adc32c3fe6434dc31cb1aaf3ba4102))
* **k8s:** cache resource lists via shared informer (closes [#86](https://github.com/rswanson/lfk/issues/86)) ([c8578cc](https://github.com/rswanson/lfk/commit/c8578cc3fb4b7997a780741ddbe14e54b6e807fe))
* **k8s:** surface ephemeral containers in pod views ([#180](https://github.com/rswanson/lfk/issues/180)) ([ac1a1c5](https://github.com/rswanson/lfk/commit/ac1a1c54baf293482a7a29666336816838713332))
* **karpenter:** first-class actions for NodePool / NodeClaim / EC2NodeClass ([#223](https://github.com/rswanson/lfk/issues/223)) ([5f37b70](https://github.com/rswanson/lfk/commit/5f37b70e21a4edf09aee9416332784e41ed3c15a))
* **knative:** first-class Knative Serving (Activate) + Eventing icons ([#224](https://github.com/rswanson/lfk/issues/224)) ([e89be5b](https://github.com/rswanson/lfk/commit/e89be5baa1405e366334a52bb999e494aad63e74))
* **localcluster:** manage kind/k3d/minikube clusters from inside lfk ([#175](https://github.com/rswanson/lfk/issues/175)) ([3c85fd9](https://github.com/rswanson/lfk/commit/3c85fd9955f5fd88dc78efefd298cac11f0e6bf3))
* **logger:** include kubeconfig path in kubectl/helm command logs ([c0c226c](https://github.com/rswanson/lfk/commit/c0c226c04a403d4588f4f272842c22c177abd359))
* **logger:** redact secrets, log mutation intent, surface silent errors ([c56e1ff](https://github.com/rswanson/lfk/commit/c56e1ff1d5741fb6eedf9c2a37a24d27ca506bec))
* **logger:** surface silent failures with dedup to in-app log ([#268](https://github.com/rswanson/lfk/issues/268)) ([ccfece0](https://github.com/rswanson/lfk/commit/ccfece0682b69b165c21edae994320d8440bba25))
* **logs:** add java (spring boot/logback) and postgresql preview formatters ([a5802b3](https://github.com/rswanson/lfk/commit/a5802b325cc35784828445661aacaf64fc54d63c))
* **logs:** add klog preview formatter ([f94ad6c](https://github.com/rswanson/lfk/commit/f94ad6c449171997aebcbe7c03d84ab4ddfcb74f))
* **logs:** add nginx/apache and envoy access log preview formatters ([5564e47](https://github.com/rswanson/lfk/commit/5564e47b1115b0b0f7004108c815171dccd38e42))
* **logs:** add zap dev encoder preview formatter ([1a43153](https://github.com/rswanson/lfk/commit/1a431532bbd49f3525ca98062f251e39a6d7bfdd))
* **logs:** copy save path to clipboard, log it via slog ([ca342dd](https://github.com/rswanson/lfk/commit/ca342ddaeec6513067d1f8aa34008acf1f29ba44)), closes [#61](https://github.com/rswanson/lfk/issues/61)
* **logs:** J/K scroll the structured preview side panel ([8132555](https://github.com/rswanson/lfk/commit/8132555dd6c23fbf74f59724ded4e61d83f5b2ee))
* **logs:** persistent search history with Up/Down recall ([58d6b08](https://github.com/rswanson/lfk/commit/58d6b08693996fb7d292b471de5159d424133119))
* **logs:** persistent search history with Up/Down recall in log viewer ([cc70537](https://github.com/rswanson/lfk/commit/cc7053710e8f338aa7bbc70b66ae32d8f4f5c5d6))
* make it possible to load namespace with bookmarks using &lt;tab&gt; ([45a9a8a](https://github.com/rswanson/lfk/commit/45a9a8a0bf5e6c908008a508edf2ed2790a35202))
* mouse wheel scrolls in YAML, Describe, Diff, Help, Explain modes [#42](https://github.com/rswanson/lfk/issues/42) ([92b5c51](https://github.com/rswanson/lfk/commit/92b5c51062f5dfe53d4f21815c1346e8ce4a8681))
* **mouse:** click-to-drill, right-click action menu, overlay mouse ([8287ba0](https://github.com/rswanson/lfk/commit/8287ba0b3fa5c50f462f063293c5915f3871a51c))
* namespace selector A binding (and only A) flips to all-namespaces ([f289b06](https://github.com/rswanson/lfk/commit/f289b06fb8ed3af97c9eb512238ef1e9a7b227b5))
* **networking:** per-endpoint preview for Endpoints / EndpointSlices ([fb0201d](https://github.com/rswanson/lfk/commit/fb0201d9513a19ff9c3ad429a0a331b5916f9af5))
* **networking:** Service preview rollup of backing EndpointSlices ([66a1e26](https://github.com/rswanson/lfk/commit/66a1e26b74894eafd53b746012bd8eeb50586c40))
* **nodes:** format CPU/Mem alloc in node preview to human-readable units ([026066e](https://github.com/rswanson/lfk/commit/026066edb628bd4a556e7ccf1688ac80c6016c56))
* **nodes:** rename Alloc → Avail and format values consistently ([3939411](https://github.com/rswanson/lfk/commit/3939411b0d482e082bbc427acdd75cbce0c2bc20))
* **palette:** add :errors / :bookmarks / :reload commands ([fcd4fd4](https://github.com/rswanson/lfk/commit/fcd4fd409d41f7b775776b41134fec604cf9321c))
* per-kind sort memory and per-context column memory ([#310](https://github.com/rswanson/lfk/issues/310)) ([4e80556](https://github.com/rswanson/lfk/commit/4e80556ba5e10368ac63804ef5798c80530259bf))
* **perf:** add 1Hz runtime sampler with ring buffer ([d576350](https://github.com/rswanson/lfk/commit/d576350a3d58d460fe53b622eee27c06a1d3c628))
* **perf:** add active-scripted scenario with key script ([0b6f6c6](https://github.com/rswanson/lfk/commit/0b6f6c6680a7a30022152719d1a12487804deeb0))
* **perf:** add deterministic small/medium/large fixtures ([c79f41c](https://github.com/rswanson/lfk/commit/c79f41c371f72920c33ad6e112353fb288608107))
* **perf:** add energy probe lifecycle skeleton ([d243e9c](https://github.com/rswanson/lfk/commit/d243e9c8311cd13798b0d32b9dae66701d41d7de))
* **perf:** add idle-background scenario using osascript ([c553b7f](https://github.com/rswanson/lfk/commit/c553b7ff6a59b467ecfbe61b3097a1f66f7ed889))
* **perf:** add labelled energy.Tick wrapper ([e015b0e](https://github.com/rswanson/lfk/commit/e015b0ed57a4f078ad6835d635ee92e799b0448e))
* **perf:** add minimal fake apiserver for energy benchmarks ([b3a7500](https://github.com/rswanson/lfk/commit/b3a7500ab606d6c199eeec02b93278cc18a0224e))
* **perf:** add powermetrics sampler for P/E-core residency ([a1f0208](https://github.com/rswanson/lfk/commit/a1f02086b53e66f7893489b72c1896a434d8d7c3))
* **perf:** add PTY driver and idle-foreground scenario ([7c030fc](https://github.com/rswanson/lfk/commit/7c030fcbf8151f18e466e6163260d3b3cd11e9a4))
* **perf:** add top(1) sampler for CPU and wakeups ([e923b76](https://github.com/rswanson/lfk/commit/e923b76390da6ab973255b3763f596eadfa573c9))
* **perf:** aggregate reports with baseline diff and markdown output ([65e7a8a](https://github.com/rswanson/lfk/commit/65e7a8a41503dc6ef6c9e4aac261e9a636ae3e13))
* **perf:** flush probe samples to JSONL on shutdown ([614147e](https://github.com/rswanson/lfk/commit/614147e21f9e4651b28dea24c9ceb2ba6e630095))
* **perf:** make the blurred/idle watch interval configurable ([#8](https://github.com/rswanson/lfk/issues/8)) ([2f37a38](https://github.com/rswanson/lfk/commit/2f37a3837e7e06c3cf18441472aef7f0af2cc852))
* **perf:** route informer + probe goroutines to E-cores via macOS QoS ([#7](https://github.com/rswanson/lfk/issues/7)) ([b05c985](https://github.com/rswanson/lfk/commit/b05c985dd5f9b3a21e3aa1cd8e80f3c88e192355))
* **perf:** scaffold energy-bench driver with kubeconfig writer ([3165c51](https://github.com/rswanson/lfk/commit/3165c51956171f98bafeb9362d932c65b5ea53a3))
* **perf:** start energy probe from main and flush on SIGUSR1 ([79c95d8](https://github.com/rswanson/lfk/commit/79c95d8b1c268954c1cd73c6f890162275675e0e))
* **perf:** wire energy-bench end-to-end with report output ([a1aaaed](https://github.com/rswanson/lfk/commit/a1aaaed610715d892c7ce5ea6f19308bbd80b37f))
* persistent up/down history for / and f search/filter ([3efa1ad](https://github.com/rswanson/lfk/commit/3efa1adb6b9d88cfd04dcf77c8bbe3935cf2e804)), closes [#54](https://github.com/rswanson/lfk/issues/54)
* pin individual resource types into a top "Pinned" section ([#300](https://github.com/rswanson/lfk/issues/300)) ([3953bd6](https://github.com/rswanson/lfk/commit/3953bd63416313413705a420b19e8daf32a9b3d2))
* **rbac:** reverse-RBAC "Who-Can" view, layered on the Can-I overlay ([7598c68](https://github.com/rswanson/lfk/commit/7598c68501a0db5067614c8cbf35f3e91020d7fc))
* refresh namespace completion cache every 60s ([19b9541](https://github.com/rswanson/lfk/commit/19b9541d52d71630c65f091d5d965bb945afa7f2))
* refresh namespaces in the namespace selector with R ([#292](https://github.com/rswanson/lfk/issues/292)) ([bc6ace6](https://github.com/rswanson/lfk/commit/bc6ace6aaf876eda1a5c48260e909a2c421a2db6))
* **release:** add AUR channel (lfk-bin) ([#174](https://github.com/rswanson/lfk/issues/174)) ([c6df49d](https://github.com/rswanson/lfk/commit/c6df49ddd153d0d15148e0dfa181f9440a19a702))
* **release:** add cloudsmith deb+rpm channel ([#163](https://github.com/rswanson/lfk/issues/163)) ([8c50bec](https://github.com/rswanson/lfk/commit/8c50bec5d2e7fb158d960543939b57b171b2915a))
* **release:** add scoop, winget, chocolatey channels ([#161](https://github.com/rswanson/lfk/issues/161)) ([f4fe4a5](https://github.com/rswanson/lfk/commit/f4fe4a576486b3c84e1949f773e6cfc3bad0ad84))
* **release:** foundation for new package-manager channels ([#159](https://github.com/rswanson/lfk/issues/159)) ([0a4a353](https://github.com/rswanson/lfk/commit/0a4a353a3e715b05a6b5b86e2a4183e26c9b21e9))
* render ansi colors, make it possible to disable ansi rendering ([3b88b82](https://github.com/rswanson/lfk/commit/3b88b8227b91a1ccc36ed059ed7d70d5450d58ba))
* render right column details even if there's no child resources ([bc8c75f](https://github.com/rswanson/lfk/commit/bc8c75f1c21a4f100721470e9c916378552db8cc))
* **resource-map:** traverse Pod refs with MissingRef detection ([a0517c3](https://github.com/rswanson/lfk/commit/a0517c306964e02bfd8a3126688727ec4f96bf77))
* **scheduler:** priority task queue with per-context dispatch ([#186](https://github.com/rswanson/lfk/issues/186)) ([80e0ba1](https://github.com/rswanson/lfk/commit/80e0ba13f5dc7d4a3609e3233002752c6686d8c6))
* scroll overflow indicator + namespace overlay no-match copy fix ([425f00f](https://github.com/rswanson/lfk/commit/425f00f8a8b62150be510eb3de746358b47dce5a))
* **security:** security findings dashboard ([#183](https://github.com/rswanson/lfk/issues/183)) ([2e3c0d9](https://github.com/rswanson/lfk/commit/2e3c0d968f86977436bc08368b06e9959cbca268))
* send arrow key sequences in DECCKM application in exec view ([9bbcb5c](https://github.com/rswanson/lfk/commit/9bbcb5c026cc1d134cebc864d9375ec0b414dece))
* shift+r at LevelResourceTypes re-runs API discovery ([4246482](https://github.com/rswanson/lfk/commit/4246482da3f090c490ad3576e24ed938264552af))
* show cursor block in exec terminal display [#24](https://github.com/rswanson/lfk/issues/24) ([e0f4d43](https://github.com/rswanson/lfk/commit/e0f4d433544007de6aa202706118312bd55ed5ca))
* show nodeport for nodeport services [#46](https://github.com/rswanson/lfk/issues/46) ([7721af1](https://github.com/rswanson/lfk/commit/7721af12811324d3b46c1c7b1d9a05d15f5564ee))
* support PgUp/PgDown/Home/End navigation keys [#35](https://github.com/rswanson/lfk/issues/35) ([c633055](https://github.com/rswanson/lfk/commit/c6330550d973b9a392349183203a706b07272355))
* Tab toggles broad mode in / and f to also match column values [#43](https://github.com/rswanson/lfk/issues/43) ([3010bb9](https://github.com/rswanson/lfk/commit/3010bb9b3a8fed49c5892633765afaae4bb1f86a))
* traffic capture (kubectl-debug + kubeshark backends) ([#179](https://github.com/rswanson/lfk/issues/179)) ([b51d64c](https://github.com/rswanson/lfk/commit/b51d64c421a4001ff1cf1c53efc2792cceb41b3f))
* **ui:** abbreviate long pod statuses when layout is too narrow ([f146d55](https://github.com/rswanson/lfk/commit/f146d55683e2644c41e24241695920c6f14b4864))
* **ui:** add structured preview side panel to log viewer ([b7074dd](https://github.com/rswanson/lfk/commit/b7074dde12d76d806efd8b749aaf84ab3f81a38b))
* **ui:** add y to copy cursor row from rollback / history overlays ([5cc7cf6](https://github.com/rswanson/lfk/commit/5cc7cf64009ac20f9e826d2373274479c5f717d9))
* **ui:** advertise y/n alongside Enter/Esc for confirm dialogs ([4f95b5c](https://github.com/rswanson/lfk/commit/4f95b5cf4140371c1cfa00fb4cf78b1faca9c779))
* **ui:** auto-apply single result in colorscheme selector ([1e944c8](https://github.com/rswanson/lfk/commit/1e944c8b77265840a7aff7bc15d049207f9a2a39))
* **ui:** auto-apply single result in container selector ([37f6a8c](https://github.com/rswanson/lfk/commit/37f6a8c11b77cefdbe9bc7730c4a378dc1cfacd6))
* **ui:** auto-apply single result in template selector ([4eed898](https://github.com/rswanson/lfk/commit/4eed8985011d4b1659344148b336bd5b74badbfc))
* **ui:** dim explorer behind overlays via dim_overlay option ([#99](https://github.com/rswanson/lfk/issues/99)) ([df167f4](https://github.com/rswanson/lfk/commit/df167f4bb20f65eda38b72ff80424755d8deb8d5))
* **ui:** migrate HelmHistory / HelmRollback / DeploymentRollback to OverlayList ([#234](https://github.com/rswanson/lfk/issues/234)) ([1fbe874](https://github.com/rswanson/lfk/commit/1fbe87495036b72b7c9ec1646810911086f81450))
* **ui:** pin info chips far-right + entry-aware keymap fit ([#101](https://github.com/rswanson/lfk/issues/101)) ([1163c7a](https://github.com/rswanson/lfk/commit/1163c7a9ea928e484d75b500f7db3029035ac029))
* **ui:** shorten verbose column headers with display aliases ([ab0b92a](https://github.com/rswanson/lfk/commit/ab0b92ad21d9b69d616cd1c72342b85c2fcd4218))
* **ui:** support count-prefixed motion (Nj/Nk) in read-only viewers ([1068839](https://github.com/rswanson/lfk/commit/10688392fb6bf4b49d6a8b20bd3ada5ad1b3335f))
* **ui:** support count-prefixed yank (Ny) in read-only viewers ([28782d1](https://github.com/rswanson/lfk/commit/28782d1e9064fe83339192d8f406e7d93b516c9a))
* **ui:** unified overlay components — OverlayList, OverlayConfirm, OverlayInput ([#231](https://github.com/rswanson/lfk/issues/231)) ([5960090](https://github.com/rswanson/lfk/commit/59600903183f6a7891282e3c23e335567c6b4362))
* unify y/n confirm overlays on Enter/Esc and polish quit dialog [#45](https://github.com/rswanson/lfk/issues/45) ([61790cf](https://github.com/rswanson/lfk/commit/61790cfddb22057f7608ffbca4f81e0b096cc458))
* **viewers:** extend count-prefix to column / word / page / search motions ([64e9498](https://github.com/rswanson/lfk/commit/64e9498894acd84684a57679b9387c364de36f56))
* **viewers:** match vim/nvim [count]&lt;C-d&gt;/&lt;C-u&gt; 'scroll' option semantics ([fd83a57](https://github.com/rswanson/lfk/commit/fd83a57cd57c5bc950d507ad78ce7c38f0e00bec))
* **viewers:** vim text-object selection (viw/vaw/viW/vaW) ([#185](https://github.com/rswanson/lfk/issues/185)) ([7eb0aea](https://github.com/rswanson/lfk/commit/7eb0aeaf2b5263e7c7fd630b3ac300b7e5140061))
* warm namespace cache on context open, invalidate on mutation ([19b9541](https://github.com/rswanson/lfk/commit/19b9541d52d71630c65f091d5d965bb945afa7f2))


### Bug Fixes

* **actions:** block delete keypress in containers view ([#181](https://github.com/rswanson/lfk/issues/181)) ([584ff4f](https://github.com/rswanson/lfk/commit/584ff4fd0bf50f3ba1c65b061469bcf741e09804))
* add gateways and tlsroutes to the networking category [#33](https://github.com/rswanson/lfk/issues/33) ([cdf1cd6](https://github.com/rswanson/lfk/commit/cdf1cd6e7713d15a44a77766a9abd80e0079469b))
* address CodeRabbit findings on PR [#122](https://github.com/rswanson/lfk/issues/122) ([1a0a97c](https://github.com/rswanson/lfk/commit/1a0a97cedafa22d2f623b1cb5091c73f2fd6d462))
* address second round of CodeRabbit findings on PR [#122](https://github.com/rswanson/lfk/issues/122) ([3fffdd4](https://github.com/rswanson/lfk/commit/3fffdd4ff336da514768307b131c79067e696252))
* also cancel log streams on :q/:quit; extract cancelActiveTabLogStreams ([bd791d0](https://github.com/rswanson/lfk/commit/bd791d02b52e93e97a11a4c3241314b70ed00c4f))
* also cancel log streams on :q/:quit; extract cancelActiveTabLogStreams ([31dc32e](https://github.com/rswanson/lfk/commit/31dc32ea9794396062fde6a2f2e4e2d07d266591))
* **app,scheduler:** drop idle CPU from ~145% to ~0% (closes [#206](https://github.com/rswanson/lfk/issues/206)) ([#211](https://github.com/rswanson/lfk/issues/211)) ([bcdd3d0](https://github.com/rswanson/lfk/commit/bcdd3d0ce10cb924a4e58703cc64367ddb7c5688))
* **app,ui:** broad-mode search/filter cycles all matched-group members; gate bar highlight on Tab ([59f4a19](https://github.com/rswanson/lfk/commit/59f4a19c91f0ddb1b9e9b76114b17c536c3b2c0b))
* **app:** apply Y bulk to LevelOwned and skip false bulk at LevelContainers ([e366579](https://github.com/rswanson/lfk/commit/e366579924f2d286702f7d6d0c9bf49f2826e949))
* **app:** clear search highlight on level-change navigation ([00466d5](https://github.com/rswanson/lfk/commit/00466d5ee4ddbb626f62f569d3050a89977615c9))
* **app:** gate category matching for both / and f on Tab (broad mode) ([0674e07](https://github.com/rswanson/lfk/commit/0674e07eab083bd5ce1e87cf5bafcd664eef8050))
* **app:** keep silent ns refresh from clobbering an open overlay ([73d5ba2](https://github.com/rswanson/lfk/commit/73d5ba2dd8483ddbc2980559ca1a06306097ea5f))
* apply theme switch to cached previews immediately ([#299](https://github.com/rswanson/lfk/issues/299)) ([7d32aaa](https://github.com/rswanson/lfk/commit/7d32aaa028d3fcc8903d962edb251cdce050076b))
* **app:** make node shell work on SELinux-enforcing immutable distros ([4477e9c](https://github.com/rswanson/lfk/commit/4477e9cc516d269dc23093c9af8c957a9c2a1689))
* **app:** plug read-only bypasses across labels and overlays ([66a11d0](https://github.com/rswanson/lfk/commit/66a11d0f421d9cde78b874e04402299b373df3f2))
* **app:** refresh right-pane preview at LevelResourceTypes on tab switch and watch tick ([#216](https://github.com/rswanson/lfk/issues/216)) ([cc4c90f](https://github.com/rswanson/lfk/commit/cc4c90f68d5e4d73c6e8915bf340fd61335c73b8))
* **app:** require typed confirmation for action-menu Force Delete ([8b1b2b7](https://github.com/rswanson/lfk/commit/8b1b2b768f37fc23d3bf0fab1143fe11b7e15a81)), closes [#89](https://github.com/rswanson/lfk/issues/89)
* **app:** respect KUBE_EDITOR and parse editor flags ([#226](https://github.com/rswanson/lfk/issues/226)) ([d944265](https://github.com/rswanson/lfk/commit/d9442656a5be06685742a63596bfd867ba879fb4))
* **app:** search no longer multiplies hits by matched category at LevelResourceTypes ([8056008](https://github.com/rswanson/lfk/commit/805600812435f30e4c5f1684486051575a64372e))
* **app:** send correct arrow key sequences in DECCKM application cursor mode ([cd02fc9](https://github.com/rswanson/lfk/commit/cd02fc9d6184d0df5bb97dc85c5c73a64649d4a0))
* **app:** two-pass search — names first, fall back to first-of-category ([4721242](https://github.com/rswanson/lfk/commit/47212420375a7866517419fe782f252d0e9e605d))
* **app:** unify Force Delete help text across menus, dialogs, and docs ([41b3aba](https://github.com/rswanson/lfk/commit/41b3abac191aa8ad7cc50b1efe37e83257815219))
* **bookmarks:** wait for discovery before declaring resource type missing ([6d6e50a](https://github.com/rswanson/lfk/commit/6d6e50a5049e597c365f9281c30e9bc1a2fdb86e))
* broad-mode persists past Enter so n/N and the applied filter keep matching ([484015f](https://github.com/rswanson/lfk/commit/484015f3bc3560faa5593700d82df83dcbe3b897))
* bulk item deletion on grouped resources ([855475c](https://github.com/rswanson/lfk/commit/855475c9fae29ab8f026277704f4f3b832eb2d35))
* **bulk:** clear stale bulk-action snapshot on dispatch and cancel ([#257](https://github.com/rswanson/lfk/issues/257)) ([16a4c60](https://github.com/rswanson/lfk/commit/16a4c606b632785b84f383251d88b31d12fd7a47))
* cancel log streams on tab close and quit ([57ed5b9](https://github.com/rswanson/lfk/commit/57ed5b96d345f8083a16856baf0b6ef71c9aa42d)), closes [#48](https://github.com/rswanson/lfk/issues/48)
* cancel log streams on tab close and quit ([e22d52e](https://github.com/rswanson/lfk/commit/e22d52e210f44fb77751826e4e34bb124aabd811)), closes [#48](https://github.com/rswanson/lfk/issues/48)
* clamp help-screen scroll so ctrl+u responds on first press ([0fdc5e5](https://github.com/rswanson/lfk/commit/0fdc5e5bfdccc609626b3b4ba4c390bf34d4057d))
* clarify rollback action descriptions to reflect picker behavior ([894d251](https://github.com/rswanson/lfk/commit/894d251b1f1875cfd414de920f6f566d9fc638f4))
* clarify rollback action descriptions to reflect picker behavior ([27243b1](https://github.com/rswanson/lfk/commit/27243b18901305b9b65aa512f44993b32e311970))
* clear hint-bar message immediately on explorer navigation ([#291](https://github.com/rswanson/lfk/issues/291)) ([3b5bb3e](https://github.com/rswanson/lfk/commit/3b5bb3e628cb1aaa55b7de899250372174d707ae))
* clear previewLoading when drilling into a pod with containers [#34](https://github.com/rswanson/lfk/issues/34) ([4a474d5](https://github.com/rswanson/lfk/commit/4a474d59d1c0775e698ef59c43d6753021abf9bc))
* clear stale pod metrics when metrics-server payload is empty ([0191775](https://github.com/rswanson/lfk/commit/01917750e54e90baba1e1694c4dd301f8f39d50a))
* clear YAML 'Loading...' placeholder when fetch errors or is canceled [#34](https://github.com/rswanson/lfk/issues/34) ([aa313cd](https://github.com/rswanson/lfk/commit/aa313cd99ede8c14fee9b3e232ed57e8e392aab3))
* cluster dashboard rendering, scrolling, and responsive usage bars ([#293](https://github.com/rswanson/lfk/issues/293)) ([#294](https://github.com/rswanson/lfk/issues/294)) ([9b3fb00](https://github.com/rswanson/lfk/commit/9b3fb001d4a32f2a7e51780f5dda4e1a1cffae43))
* column toggle overlay box stays stable size while filtering ([dbc4760](https://github.com/rswanson/lfk/commit/dbc47602a248945892ac4a3c92d5903237af9f50))
* column toggle overlay filter bar anchored under title (matches namespace overlay) ([6769149](https://github.com/rswanson/lfk/commit/67691497fc2cc3ef437c47ac92fa315d4bb6e663))
* **commandbar:** -A flag, pty inline output, exact-match autocomplete ([#235](https://github.com/rswanson/lfk/issues/235)) ([040a690](https://github.com/rswanson/lfk/commit/040a6901b17616cb7df3adb0ab98652111eef912))
* **copy:** JSON shortcut chip + partial-success on bulk YAML/JSON copy ([#239](https://github.com/rswanson/lfk/issues/239)) ([74ec4df](https://github.com/rswanson/lfk/commit/74ec4df9e173f6a30764015d89ffac765c3f4f75))
* **cursor:** preserve syntax/highlight styling on cursor row ([b6f55a2](https://github.com/rswanson/lfk/commit/b6f55a2a35ea28346079a23a76a0a5bc28357796))
* dashboard events preview background tear under non-black themes ([#293](https://github.com/rswanson/lfk/issues/293)) ([#296](https://github.com/rswanson/lfk/issues/296)) ([2d1672d](https://github.com/rswanson/lfk/commit/2d1672dc7272759cf72904a93982f4c4d4fc644a))
* default to UTC when spec.timeZone is empty; tidy go.mod; doc + test polish ([00adcb5](https://github.com/rswanson/lfk/commit/00adcb50ca3f041055afd11b500b3debdc767bfe))
* **describe:** route keys to search input, not global tab handler ([#203](https://github.com/rswanson/lfk/issues/203)) ([#204](https://github.com/rswanson/lfk/issues/204)) ([40ea18d](https://github.com/rswanson/lfk/commit/40ea18de1aeb7ad4dc3c5ee29573da9688d7834d))
* details pane shows children + theme color tracking ([#238](https://github.com/rswanson/lfk/issues/238)) ([61e0ca3](https://github.com/rswanson/lfk/commit/61e0ca3f6a42213352d715180cf28cd071ca5e8c))
* discovery completion uses middleItems-empty (not m.loading) as the ([a25a1f6](https://github.com/rswanson/lfk/commit/a25a1f6434a3332900b1fef7ae95e2c0b264ac00))
* don't bleed search highlight into the parent (left) column ([b7a3eaa](https://github.com/rswanson/lfk/commit/b7a3eaa67dd921e4144dedbd8aaff0f43fa7cb79))
* don't include prefixes and timestamps when they are hidden [#40](https://github.com/rswanson/lfk/issues/40) ([37e7a5f](https://github.com/rswanson/lfk/commit/37e7a5f47d3a13429e00bd68a4cc9b4efa877eb2))
* **editors:** address CodeRabbit review on PR [#134](https://github.com/rswanson/lfk/issues/134) ([c18c1ce](https://github.com/rswanson/lfk/commit/c18c1ce1debc0296747f146282803c7d9af3ff88))
* **editors:** ANSI leak in field labels + up/down nav + scroll-to-cursor ([b82144c](https://github.com/rswanson/lfk/commit/b82144c9a22b0e647ba984d9aa66bcec76f440ce))
* **editors:** collapse long/multi-line values to a single visual cell ([97481d5](https://github.com/rswanson/lfk/commit/97481d5aa6ebf98798a2e12caba335a081f409f9))
* **editors:** consistent key column + space-select + smart-y ([1302e3b](https://github.com/rswanson/lfk/commit/1302e3b46cb8e31149f109080298faba9adb9c13))
* **editors:** ctrl+s under active filter no longer mutates wrong key ([c02de2b](https://github.com/rswanson/lfk/commit/c02de2b35d6bf09ee4a538aa1bc5726ebbd01663))
* **editors:** cursor in edit pane lands at TextInput cursor pos + ([6ab3526](https://github.com/rswanson/lfk/commit/6ab35266d90fa34f6a1d82b53d32ee91269d5b5c))
* **editors:** format picker no longer shrinks the table ([1e0a2eb](https://github.com/rswanson/lfk/commit/1e0a2eb1d087a0e8b25fe39b5d48b501c498a50d))
* **editors:** inline edit mode for single-line values ([2cc3e53](https://github.com/rswanson/lfk/commit/2cc3e53be4b80e377ff9903510d68b35b63d368b))
* **editors:** show multi-line values as multi-line during editing ([6c59ceb](https://github.com/rswanson/lfk/commit/6c59ceb9c4d79ef96d72655fd250e5fca10b1e19))
* **editors:** sticky scroll + ctrl+u/d/f/b page keys + line-scoped ctrl+a/e ([1d8ba02](https://github.com/rswanson/lfk/commit/1d8ba021ab7ba45db07bbc1a4ea91235b2f0e313))
* **exec,browser:** make interactive shell + browser-open actions work on Windows ([#197](https://github.com/rswanson/lfk/issues/197)) ([4954439](https://github.com/rswanson/lfk/commit/4954439941e12860973643a0f6a91b36232a713e))
* explorer Esc clears search highlights before navigating parent ([b7a3eaa](https://github.com/rswanson/lfk/commit/b7a3eaa67dd921e4144dedbd8aaff0f43fa7cb79))
* **filter:** clear active filter preset on Esc ([#156](https://github.com/rswanson/lfk/issues/156)) ([7b22dff](https://github.com/rswanson/lfk/commit/7b22dff6d59d31d8484536abe445cd6774be45a8))
* **filter:** clear stale preview when filter preset matches zero items ([#157](https://github.com/rswanson/lfk/issues/157)) ([5a8f28c](https://github.com/rswanson/lfk/commit/5a8f28c79aa2579ab184a42b8184855f059469f9))
* follow kubeconfig.d symlinks [#23](https://github.com/rswanson/lfk/issues/23) ([788bc1d](https://github.com/rswanson/lfk/commit/788bc1dbeb08e3b6ee1bac212bf85665ba144f16))
* gate explorer hint bar by nav level at context picker ([#288](https://github.com/rswanson/lfk/issues/288)) ([199de4b](https://github.com/rswanson/lfk/commit/199de4b9730e1a00e349d14b84daec7bb662c874))
* gate older-history auto-load on cursor at top, not just scroll==0 ([8fc704b](https://github.com/rswanson/lfk/commit/8fc704b3af05f49f796c9341cccbede0b84f10b6))
* hardcoded watch interval [#19](https://github.com/rswanson/lfk/issues/19) ([905a00d](https://github.com/rswanson/lfk/commit/905a00d03c7ea0c79f69b0bfd3928651fd463708))
* help screen box stays the same height when filtering ([78d2dff](https://github.com/rswanson/lfk/commit/78d2dff0d5ffc81846c0d24080aa301d57962e4c))
* help-scroll clamp uses the renderer's actual visible-rows formula ([6868a5c](https://github.com/rswanson/lfk/commit/6868a5c2a13f1045dee4c92178b5d0acec39015e))
* **help:** address lint and CodeRabbit review ([490fe6d](https://github.com/rswanson/lfk/commit/490fe6daf363b8011508a08fea8f2ab321eaf70d))
* highlight selected resource type in parent pane at LevelResources ([00f5163](https://github.com/rswanson/lfk/commit/00f51639e863c377d5c02915b59ee9172e79ff13))
* **history:** preserve draft on edit-after-recall via leaveBrowse() ([321a4bc](https://github.com/rswanson/lfk/commit/321a4bc4e75c62bc3a3110996fe8a3f687600352))
* **history:** tighten file perms and leaveBrowse on paste ([0678df4](https://github.com/rswanson/lfk/commit/0678df45075c3b136359a42db236ff1ee2f673e7))
* honor CRD additionalPrinterColumns priority and keep them visible ([#309](https://github.com/rswanson/lfk/issues/309)) ([bed3a7a](https://github.com/rswanson/lfk/commit/bed3a7af0d241c34bae675af80d0c1c93544f199))
* invalidate right-pane preview on search/filter cursor jump ([c35b4b1](https://github.com/rswanson/lfk/commit/c35b4b120b8a199f504778b9d3e356f753628cc5))
* **k8s,ui:** harden informer cache wiring + config parsing ([ddd596e](https://github.com/rswanson/lfk/commit/ddd596eb2f21f86d5d863aab50a4dc9c803ac60b))
* **k8s:** dedup kubeconfig paths so collectContexts doesn't see one file twice ([4eace3d](https://github.com/rswanson/lfk/commit/4eace3daa173a533644fff72a101a497f89b8188))
* **k8s:** isolate per-context resolution from kubeconfig merge collisions ([aebb17c](https://github.com/rswanson/lfk/commit/aebb17c578e047c0c0b0f004c42be4aeec126c23))
* key namespace completion cache by context ([f05125a](https://github.com/rswanson/lfk/commit/f05125a9a858912b335f02a732fbde59979f30a4)), closes [#29](https://github.com/rswanson/lfk/issues/29)
* **kv-editor:** dedicated selection column for the Secret/ConfigMap/Label editors ([#240](https://github.com/rswanson/lfk/issues/240)) ([2c00356](https://github.com/rswanson/lfk/commit/2c00356b486f11e5d50626dbcdf819cb4f84dca2))
* link manually-triggered CronJob jobs to their CronJob via ownerRef ([#308](https://github.com/rswanson/lfk/issues/308)) ([39acb48](https://github.com/rswanson/lfk/commit/39acb485a3391fc48491481fc09e2a348db52594)), closes [#304](https://github.com/rswanson/lfk/issues/304)
* **lint:** inline int32Ptr and use slices.Backward (golangci-lint v2.12) ([f8f47a1](https://github.com/rswanson/lfk/commit/f8f47a17fb64a4b354be7c32eae1bd48330d53f5))
* **logs:** clamp rune-slice index in log search to prevent panic ([9068026](https://github.com/rswanson/lfk/commit/9068026582c5803bf363f6a823eb4e3b838adc61))
* **logs:** defend against rows that overflow the body or contain newlines ([f6ede0a](https://github.com/rswanson/lfk/commit/f6ede0a17b5de079ce4a6f1c590daf9bce0370be))
* **logs:** expand tabs in sanitizer to fix dragonfly border push-off ([34092e1](https://github.com/rswanson/lfk/commit/34092e18f12b2ff14119e3a530ba7d9360eafc6d))
* **logs:** handle Ctrl+U (delete-line) in log viewer search input ([cd1a049](https://github.com/rswanson/lfk/commit/cd1a0495721989cb574777ccd6876242870db6f0))
* **logs:** keep cursor at top when older history loads after gg ([#241](https://github.com/rswanson/lfk/issues/241)) ([add9187](https://github.com/rswanson/lfk/commit/add918752cdcd067bc4151bbcee28bd4b0bff1c9))
* **logs:** keep visual selection bg alive across embedded ANSI ([e185ef3](https://github.com/rswanson/lfk/commit/e185ef3ee87216bad63aa341419957d468d94021))
* **logs:** make char/block visual selection visual-column- and ANSI-aware ([3eb50f9](https://github.com/rswanson/lfk/commit/3eb50f940c3c5621b0df570ed2a9c8a3c64aaf28))
* **logs:** preserve embedded SGR in cursor split for kyverno-style lines ([272153b](https://github.com/rswanson/lfk/commit/272153bb31c675677a085af9c36a0e8ca524511e))
* **logs:** scope log-search backspace reset() inside len-guard ([6608c14](https://github.com/rswanson/lfk/commit/6608c14640902b22495c82a03709c2410978cab4))
* **logs:** scroll preview J/K reaches the actual last body row ([f3786b2](https://github.com/rswanson/lfk/commit/f3786b2ff1ff7fa42396231668456c2362f7b0af))
* **logs:** strip producer ANSI inside line-mode visual selection ([8a43c26](https://github.com/rswanson/lfk/commit/8a43c268d11b4b88142a00552727652ab5b57f6b))
* **logs:** use display line for wrap math so tail stays visible ([a3d9739](https://github.com/rswanson/lfk/commit/a3d9739e946ccdad3f842cbd458766ef4cf97e00))
* **logs:** visual selection pre-trim must use visual width, not runes ([b6318bd](https://github.com/rswanson/lfk/commit/b6318bde7d33bea4e73f9ed4b7cf68f2ea21fb32))
* **logs:** wrap by visual width, preserving embedded SGR sequences ([24a6ec9](https://github.com/rswanson/lfk/commit/24a6ec9e8430f66f8bdc03981505b93f9240da07))
* **logs:** wrap-aware cursor visibility and follow-pin via topSkip ([852f1a4](https://github.com/rswanson/lfk/commit/852f1a4d08d12661ed7a79b5f287f9e5cc6d4872))
* mark pods as namespaced in GetPodYAML so Enter loads pod YAML [#34](https://github.com/rswanson/lfk/issues/34) ([38c52a2](https://github.com/rswanson/lfk/commit/38c52a25cec67148b19722d31ce503cf1003e247))
* **metrics:** fall back to metrics-api when prometheus route fails ([#266](https://github.com/rswanson/lfk/issues/266)) ([3a79c37](https://github.com/rswanson/lfk/commit/3a79c377b9abb45019e3029066fe2cee840c4518))
* **metrics:** stop ~1Hz column-order blink on PodInitializing rows ([0895f56](https://github.com/rswanson/lfk/commit/0895f563667f9ae1cd25a08d2f0ac33a5a21f111))
* **metrics:** stop ~1Hz column-order blink on PodInitializing rows ([b1b53cf](https://github.com/rswanson/lfk/commit/b1b53cf05b1bc8fbddd0188b80d9727f162d3f76))
* **mouse:** address CodeRabbit review on PR [#135](https://github.com/rswanson/lfk/issues/135) ([ed9293b](https://github.com/rswanson/lfk/commit/ed9293b2e3c9741e8afff240be802d253db74ad9))
* **nav:** clear filter state when navigating to parent ([9977274](https://github.com/rswanson/lfk/commit/99772748fba84eb64889d63bd417dd0f597cf007))
* **nav:** preserve cursor on watch-tick discovery failure at LevelResourceTypes ([057f036](https://github.com/rswanson/lfk/commit/057f036fb904e2df8494ebdf65b1cb65796745b1))
* **nav:** stop using Esc to walk back through navigation levels ([1ec70a3](https://github.com/rswanson/lfk/commit/1ec70a3c58bad0dadaf00d926ef31aa1d8e2880d))
* **networking:** address coderabbit findings on Service endpoints rollup ([de75cb3](https://github.com/rswanson/lfk/commit/de75cb31dfdfa87d1134ddc05ab9b64064680293))
* **networking:** always refetch Service endpoints; cache hid pod churn ([906b37f](https://github.com/rswanson/lfk/commit/906b37fb80b547be611f75203e8546352acfea4e))
* **networking:** carry over Service rollup columns across watch-tick rebuilds ([e67c272](https://github.com/rswanson/lfk/commit/e67c272da3f17da5b16b35bb2606dea8600882ed))
* **networking:** stale-while-revalidate Service endpoints to stop the flash ([24b6f95](https://github.com/rswanson/lfk/commit/24b6f95e5b78806787693c3150883b57168f3472))
* **networking:** treat absent EndpointSlice conditions.ready as ready ([94ac9fc](https://github.com/rswanson/lfk/commit/94ac9fc08c66a5691ff68df90b33691a49b926e9))
* **nix:** build with Go 1.26.3 by overriding pkgs.go_1_26 ([#228](https://github.com/rswanson/lfk/issues/228)) ([f4d0ad6](https://github.com/rswanson/lfk/commit/f4d0ad6c3b8b527af26171fca757886f4ff72d21))
* **nix:** update vendorHash after k8s dependency bump ([#285](https://github.com/rswanson/lfk/issues/285)) ([9020c87](https://github.com/rswanson/lfk/commit/9020c87fe35b4d8b60cfae0c23ab8ed9346bef5d)), closes [#284](https://github.com/rswanson/lfk/issues/284)
* **nodeshell:** land on DiskPressure/MemoryPressure/PIDPressure nodes ([#177](https://github.com/rswanson/lfk/issues/177)) ([eec8d02](https://github.com/rswanson/lfk/commit/eec8d02a5426c1667f9d0a195bfa8765da1c98d2))
* **nodes:** keep CPU/MEM/alloc columns visible across metrics churn ([5ccba28](https://github.com/rswanson/lfk/commit/5ccba286ae2179c2eff45458367eeccae8695821))
* normalize clipboard line endings to CRLF on Windows ([#261](https://github.com/rswanson/lfk/issues/261)) ([#265](https://github.com/rswanson/lfk/issues/265)) ([2b884aa](https://github.com/rswanson/lfk/commit/2b884aa610a49b90bc3b08ed9ce29919e5d2e2d1))
* **perf:** address final code-review issues from Phase 1 ([87faf77](https://github.com/rswanson/lfk/commit/87faf7791ca329d28c5f4b18b7d101a44638195c))
* **perf:** flush probe JSONL on PTY teardown + rebaseline ([#3](https://github.com/rswanson/lfk/issues/3)) ([5dc4d56](https://github.com/rswanson/lfk/commit/5dc4d563256229e0173b1da8d83c53806d66f8fb))
* **perf:** make sampler race-free and Stop synchronous ([77d2ef9](https://github.com/rswanson/lfk/commit/77d2ef9cf3785616b6d7bcfa518d6991b066da7a))
* **perf:** retire superseded watch-tick chains via generation token ([854ddf0](https://github.com/rswanson/lfk/commit/854ddf0b55d5a5166e6858fa0b8458694f686cd5))
* **perf:** use idlew (idle wakeups) and power columns in top sampler ([88a4324](https://github.com/rswanson/lfk/commit/88a43240ed122bcbca39793232cf63104eef4155))
* persist in-context read-only toggle to per-context override ([#290](https://github.com/rswanson/lfk/issues/290)) ([35cbab5](https://github.com/rswanson/lfk/commit/35cbab5dfdc868303c7300e8a6359a91dc60bb20))
* pod metrics enrichment in single-namespace mode and across ticks ([aa44549](https://github.com/rswanson/lfk/commit/aa4454911f8c5ab31f597ce0bd935710e576eecf))
* polish cluster-list discovery UX and auto-reconnect pod logs across container transitions ([0a881b7](https://github.com/rswanson/lfk/commit/0a881b703cbb29d1dc3108febebf828ce2068ad4))
* **port-forward:** reuse local port on restart, fix setup-overlay UX ([#253](https://github.com/rswanson/lfk/issues/253)) ([d065a5e](https://github.com/rswanson/lfk/commit/d065a5e15204f9edfb55314a8186420513eb3d59))
* preserve cursor across periodic API discovery refreshes ([e651854](https://github.com/rswanson/lfk/commit/e6518546b5cfa79fe509d71e631d8a4659e322fd))
* preserve input order and rune span in fuzzy completion ([865a20c](https://github.com/rswanson/lfk/commit/865a20c6d370889e559616df2e991c4501e21b4d))
* **preview:** clear previewLoading when resource list arrives empty ([a781377](https://github.com/rswanson/lfk/commit/a781377429081c8794ce7ccd20674b7f03cf031a))
* **preview:** DATA (N) counts keys not visual lines ([fa623d4](https://github.com/rswanson/lfk/commit/fa623d4bbff9a39749a90060f2bb4d577faf6934))
* **quit:** cancel in-flight API requests so quit doesn't hang on dead clusters ([b0479b8](https://github.com/rswanson/lfk/commit/b0479b8978a32387c15452cf0ff077e9e005664c))
* **rbac:** address remaining coderabbit findings on Who-Can ([8acdd11](https://github.com/rswanson/lfk/commit/8acdd11965492c7f1d1f6a328ad7aa1e514678b4))
* refresh actually refetches and Age advances between fetches ([cf02701](https://github.com/rswanson/lfk/commit/cf0270159bc1f21a975b82630b11dc16944eb448))
* **release:** apt install mono-devel before choco wrapper (ubuntu-24.04) ([#168](https://github.com/rswanson/lfk/issues/168)) ([28488b6](https://github.com/rswanson/lfk/commit/28488b69c73e769af194b97104f338b11c8327b4))
* **release:** declare cosign bundle as signature artifact ([#150](https://github.com/rswanson/lfk/issues/150)) ([1adf6ea](https://github.com/rswanson/lfk/commit/1adf6eae58e26f23510291c0727e60cc8b24da60))
* **release:** install chocolatey via mono wrapper per chezmoi pattern ([#166](https://github.com/rswanson/lfk/issues/166)) ([d3d0eb3](https://github.com/rswanson/lfk/commit/d3d0eb38a1ee2cb513c50214e1db2814d6133128))
* **release:** migrate cosign signing to Sigstore bundle output ([acc6284](https://github.com/rswanson/lfk/commit/acc62843638f64d1ca38a762cd1cc05295bf43e2))
* **release:** migrate cosign signing to Sigstore bundle output ([3a0c12e](https://github.com/rswanson/lfk/commit/3a0c12e8648729ede71eff55c2a0897afd353516))
* **release:** rename cosign bundle to .sigstore for Scorecard ([#152](https://github.com/rswanson/lfk/issues/152)) ([f12ce39](https://github.com/rswanson/lfk/commit/f12ce3976ceae1fd6255d573de842647a687a1a4))
* **release:** skip Chocolatey publish until first version is moderated ([#201](https://github.com/rswanson/lfk/issues/201)) ([0a5be5d](https://github.com/rswanson/lfk/commit/0a5be5d8d3fa0082576a2d04f68bf42538c6b7cc))
* **release:** use ArtifactPath template in cloudsmith publisher ([#169](https://github.com/rswanson/lfk/issues/169)) ([56b0dde](https://github.com/rswanson/lfk/commit/56b0dde59f0a7a6afa5d0065d19a4ef403ef0406))
* **release:** use pip --user for cloudsmith-cli to avoid workspace pollution ([#164](https://github.com/rswanson/lfk/issues/164)) ([1673a83](https://github.com/rswanson/lfk/commit/1673a83c0974bb5a27cf0937da3bbea763bd7498))
* remember resource list filter across subview navigation ([#303](https://github.com/rswanson/lfk/issues/303)) ([#311](https://github.com/rswanson/lfk/issues/311)) ([4dca616](https://github.com/rswanson/lfk/commit/4dca616fb12776e7e21ba3f3641e731e82e8e729))
* remove hardcoded lipgloss colored texts, use themecolor fixes [#22](https://github.com/rswanson/lfk/issues/22) ([9073a16](https://github.com/rswanson/lfk/commit/9073a16b09c2da84b7ca421148f5cdd2cf740279))
* reset query-history cursor on edits to recalled entries ([31b9909](https://github.com/rswanson/lfk/commit/31b99091cc8ff0c31f8b97285fb243c23e5596e2))
* **resource-map:** fall back to nav.Namespace at LevelContainers ([13033f3](https://github.com/rswanson/lfk/commit/13033f3195f66bb5c521f6c2aa532b467681dcad))
* **resource-map:** show Pod's tree when M is pressed at LevelContainers ([36b84fa](https://github.com/rswanson/lfk/commit/36b84fa7f81ee5dd57e386d407997f7d58f349a8))
* restrict itemCache shortcut to preview loads so deleted rows disappear ([d3094cb](https://github.com/rswanson/lfk/commit/d3094cb47f0431f9d0ebcc2b7c62e72556dbc9d7))
* **scheduler:** reclaim superseded background work so the focused view wins ([#317](https://github.com/rswanson/lfk/issues/317)) ([0e18880](https://github.com/rswanson/lfk/commit/0e188808a1a840c729fecfd53d054c3bb4196cb2))
* **scheduler:** stop title-bar spinner during 10s linger window after work completes ([#220](https://github.com/rswanson/lfk/issues/220)) ([f441c78](https://github.com/rswanson/lfk/commit/f441c7844e3bcd0450a3c1f11445633ed7879c79))
* scroll indicators split into top/bottom rows with stable layout ([eff9e4f](https://github.com/rswanson/lfk/commit/eff9e4fa1937725398f9b883012c065061690c66))
* search highlights persist past Enter; help marks current match distinctly ([b7a3eaa](https://github.com/rswanson/lfk/commit/b7a3eaa67dd921e4144dedbd8aaff0f43fa7cb79))
* search matches by category at LevelResourceTypes (with header highlight) ([b7a3eaa](https://github.com/rswanson/lfk/commit/b7a3eaa67dd921e4144dedbd8aaff0f43fa7cb79))
* search no longer matches by category, only by name (and broad-mode columns) ([b7a3eaa](https://github.com/rswanson/lfk/commit/b7a3eaa67dd921e4144dedbd8aaff0f43fa7cb79))
* **search:** paint highlight overlay live as the user types ([2a78c67](https://github.com/rswanson/lfk/commit/2a78c67c52200698d129ae3492d7d8d525c284c8))
* **session:** resume deferred CRD restore once discovery arrives ([abf5015](https://github.com/rswanson/lfk/commit/abf50158743279321b89341eaa1e616535cb4c73))
* show loader while resource types are loading ([7ee42d9](https://github.com/rswanson/lfk/commit/7ee42d9ff71a28c9b4f92948d58df8faecfda40b))
* skip non-listable resources in the sidebar ([2ed93e9](https://github.com/rswanson/lfk/commit/2ed93e97b08fea6e2fe14ede82d5c4121bfe4e4d))
* **sort:** numeric ordering for percent columns (CPU%, MEM%, */R, */L) ([#273](https://github.com/rswanson/lfk/issues/273)) ([f50537a](https://github.com/rswanson/lfk/commit/f50537a8cc91fee15fabad596f671182922a7b58)), closes [#272](https://github.com/rswanson/lfk/issues/272)
* **sort:** sort numeric/structured columns numerically ([#255](https://github.com/rswanson/lfk/issues/255)) ([1a13ad0](https://github.com/rswanson/lfk/commit/1a13ad03893d239819238cc4266765a85ef9d8b6)), closes [#250](https://github.com/rswanson/lfk/issues/250)
* stop infinite Loading spinner on permission errors ([#171](https://github.com/rswanson/lfk/issues/171)) ([07c4c14](https://github.com/rswanson/lfk/commit/07c4c14c90800144ad5999218fbfdf887069dfc9))
* stop node metrics column-order flicker (and prevent the whole class) ([#259](https://github.com/rswanson/lfk/issues/259)) ([7f5c695](https://github.com/rswanson/lfk/commit/7f5c695020bee63ed35db09ef20992a31cece50a))
* **tabs:** persist right-pane footers per tab so metrics don't bleed ([1a4fa9d](https://github.com/rswanson/lfk/commit/1a4fa9d20f3b225e3e727dd0dd72269928ad1b0f))
* **tabs:** refresh middle column on tab switch (stale-while-revalidate) ([#182](https://github.com/rswanson/lfk/issues/182)) ([b84595e](https://github.com/rswanson/lfk/commit/b84595e5718c663ed9b6096f4e0f6e3960f065d0))
* **theme:** keep parent highlight readable on themes with near-text border ([b5fc86f](https://github.com/rswanson/lfk/commit/b5fc86f4ea5b9abb0c1f54a10d356395a8f907a2))
* **ui:** auto-apply single result when committing namespace filter ([8956e63](https://github.com/rswanson/lfk/commit/8956e637efa5c00a3ded91efc425eb643ba91bfc))
* **ui:** center quit overlay text and unify confirm-hint convention ([e200ffe](https://github.com/rswanson/lfk/commit/e200ffe6cb0d140a749d695de86420dca72e5c79))
* **ui:** clear stale items when opening log container filter ([663bec1](https://github.com/rswanson/lfk/commit/663bec14e583242069b9ea7290a6fdfb918e8f58))
* **ui:** clip pinned resource-usage footer no longer triggered by event count ([#178](https://github.com/rswanson/lfk/issues/178)) ([54b6af0](https://github.com/rswanson/lfk/commit/54b6af0cdf2c56192ec15bc6233c4489cdf3753f))
* **ui:** close cache-invalidation gaps in TableRenderer ([34fe2ce](https://github.com/rswanson/lfk/commit/34fe2ce814321e8466c90ff30de6cf416fed5d7a))
* **ui:** defer log container overlay until data loads ([4e06e4a](https://github.com/rswanson/lfk/commit/4e06e4ae315632d4e4ff7b66ecaf8439d844016b))
* **ui:** drop blank line between RESOURCE USAGE header and bars ([#217](https://github.com/rswanson/lfk/issues/217)) ([40be2bc](https://github.com/rswanson/lfk/commit/40be2bc9148c41d961e1fadd0c786dfb8854a150))
* **ui:** drop dangling CONTRIBUTING.md refs and tighten confirm-hint test ([315e04d](https://github.com/rswanson/lfk/commit/315e04d07c0dbca699987f4db6ad042dfa064e59))
* **ui:** drop x:actions hint on the kubeconfig list ([c5fc303](https://github.com/rswanson/lfk/commit/c5fc303666279b7a6137d82328e688bc601ad48f))
* **ui:** filter All Containers virtual row by name in container selector ([0c879b7](https://github.com/rswanson/lfk/commit/0c879b7a6d1d2b6a91b3a4720f4edddba61df76d))
* **ui:** gate :sort command and column-header clicks on sortApplies() ([f8a7941](https://github.com/rswanson/lfk/commit/f8a794135bd5f6f8a77ae7a109aad6b862e3a9db))
* **ui:** gate log viewer n/N hint on committed search ([87f6943](https://github.com/rswanson/lfk/commit/87f694350a03289c47561515fd59ec5daf727f1a))
* **ui:** give Name column its natural width before sizing extras ([1331bd2](https://github.com/rswanson/lfk/commit/1331bd2969cf2bb9b9ba3dffc101547fe34874de))
* **ui:** give Name column natural width when room is available ([e005072](https://github.com/rswanson/lfk/commit/e005072db01a8427cb7d3a784e662ca859b1ecc6))
* **ui:** hide no-op sort and actions at picker levels ([7376449](https://github.com/rswanson/lfk/commit/7376449eb354be3449f3080a1ef1f65120f3abb6))
* **ui:** invalidate middle-column row cache on theme change ([01ca28f](https://github.com/rswanson/lfk/commit/01ca28f6963e6c0cffc9203c79a44916d67532a4))
* **ui:** keep "/" search highlight from corrupting SGR codes ([8383f6b](https://github.com/rswanson/lfk/commit/8383f6bed5fb9840f331b3510af117911c5dc4b9))
* **ui:** keep category bar underline when search highlight is active in NO_COLOR ([2a3f66d](https://github.com/rswanson/lfk/commit/2a3f66d15f7c1c920990c04d7bfbff340a396f66))
* **ui:** keep error log open behind theme selector, fix padding bg ([f789fd4](https://github.com/rswanson/lfk/commit/f789fd407edf8cebb86fee13e8eb35069e329a2f))
* **ui:** make fullscreen :errors view match other fullscreen modes ([35af38d](https://github.com/rswanson/lfk/commit/35af38de6dc58a8c96db65e916b82093faaf56c3))
* **ui:** make selected-match highlight legible across themes ([2edd86e](https://github.com/rswanson/lfk/commit/2edd86ea703a03660bca7265ed6bce55d7b4294b))
* **ui:** preserve ANSI styling when Truncate shortens a styled line ([60d5e4f](https://github.com/rswanson/lfk/commit/60d5e4f3231830d056e813084fd21874e6953429))
* **ui:** preserve level colors on the cursor line in :errors overlay ([c74cd2b](https://github.com/rswanson/lfk/commit/c74cd2beda9e099a1c271fa2a7b59ae5fec444df))
* **ui:** preserve outer background after search-highlight reset ([aecd324](https://github.com/rswanson/lfk/commit/aecd3241f5958250490d925ac3925bc1b03e0299))
* **ui:** render fullscreen errors as a viewExplorer columns slot ([0228c4c](https://github.com/rswanson/lfk/commit/0228c4cd20eb7b0e6f28be3390b28422ea2644a9))
* **ui:** route keys to overlay first, strip surfaceBg from log content ([0ae203c](https://github.com/rswanson/lfk/commit/0ae203c2c25250de52069576aa7c8eb6e49e3bf0))
* **ui:** show context-aware hint bar in YAML and log viewers ([1f35c0f](https://github.com/rswanson/lfk/commit/1f35c0f1a93928fe4164893d1130f8067a09597d))
* **ui:** show copy feedback and add y in yaml/diff/logs ([c7148af](https://github.com/rswanson/lfk/commit/c7148af3d8d3a6236a7d06209a67d831a51407d1))
* **ui:** show full hotkey hint bar with log preview on ([#71](https://github.com/rswanson/lfk/issues/71)) ([0badd03](https://github.com/rswanson/lfk/commit/0badd0302d56738c7c6934fc38b3c4f457ac4e83))
* **ui:** shrink namespace col so long pod names render without truncation [#53](https://github.com/rswanson/lfk/issues/53) ([7a9d70f](https://github.com/rswanson/lfk/commit/7a9d70f5805fd784d82f1113b7d3b1947d2faa80))
* **ui:** stop lipgloss from fragmenting embedded highlight ANSI ([197a243](https://github.com/rswanson/lfk/commit/197a243c05952506cc3e2afb062c9936e0658694))
* **ui:** stop namespace and column-toggle overlays from shrinking on filter ([7965ecd](https://github.com/rswanson/lfk/commit/7965ecd49b03690512fe2d3622a1062a3a74cec1))
* **ui:** stop selector overlays from shrinking on filter ([39d0ba2](https://github.com/rswanson/lfk/commit/39d0ba2a5a5ff952dfc2dd6bcbc170e78ced4dfa))
* **ui:** unblock tab and theme keys, fix bg in fullscreen errors view ([3bcf360](https://github.com/rswanson/lfk/commit/3bcf360ca21f86a5640c4344c04fc405c251e105))
* **ui:** widen Quick Filters overlay and clean up selected row ([86afe1a](https://github.com/rswanson/lfk/commit/86afe1acb04a3e2bb7e0fee670c931e999211a7c))
* **viewers:** clear diff digit buffer on visual mode entry ([bac6309](https://github.com/rswanson/lfk/commit/bac63095c5ae74baff9dfd0a409362ae63ae48ec))
* **viewers:** round half-page step before scaling by count ([316a91d](https://github.com/rswanson/lfk/commit/316a91dffd2446afde5c21d69f1e9b0147682c38))
* **viewers:** scale yaml page motions by viewport, not raw m.height ([67dd610](https://github.com/rswanson/lfk/commit/67dd6102ad5473f136723107fc1eed9690e4206e))
* **views:** apply GVR-keyed view columns and hide unlisted builtins ([#277](https://github.com/rswanson/lfk/issues/277)) ([a2df863](https://github.com/rswanson/lfk/commit/a2df863902a2f58f6c9dd4ed896a210d43da2051)), closes [#262](https://github.com/rswanson/lfk/issues/262)
* **yaml:** clamp rune-slice index in yaml search to prevent panic ([b501f51](https://github.com/rswanson/lfk/commit/b501f51ea9a614a1ace39c45b31743395084af06))
* **yaml:** keep syntax highlight on lines that match the search ([9245dc4](https://github.com/rswanson/lfk/commit/9245dc4f7d54cacd16da69d0f02f629db9d44248))
* **yaml:** keep token color around the search highlight ([3fc9496](https://github.com/rswanson/lfk/commit/3fc9496144af9e1ea5e7756e034687d28323ec03))


### Performance Improvements

* **app:** seed namespace selector overlay from existing cache ([d8090b4](https://github.com/rswanson/lfk/commit/d8090b42ab45cf6f485f84957d8907cd6f136a75))
* capture initial energy-bench baselines + Phase 1 polish ([44482bc](https://github.com/rswanson/lfk/commit/44482bc5836b18076cc5c4e66973fa1a1a179c00))
* **discovery:** persist API discovery to disk for stale-while-revalidate startup ([a1aaf27](https://github.com/rswanson/lfk/commit/a1aaf27622f1809ab7bef9214f3cff19e3235399))
* **explorer:** debounce preview load on cursor moves ([5d9974c](https://github.com/rswanson/lfk/commit/5d9974cf15ecec3fd8f5e0f00913e191cfeb20b9))
* **k8s:** cache API discovery to disk across sessions ([884a93b](https://github.com/rswanson/lfk/commit/884a93b99b73350e8be5ac862d1b5b58da3c2a4a))
* load PVC usage lazily via owned-children preview ([eed3e62](https://github.com/rswanson/lfk/commit/eed3e62c2e118aae1a89998db33ea9134f11e7d6))
* remove energy-bench/probe, fix inert QoS (review follow-ups) ([#10](https://github.com/rswanson/lfk/issues/10)) ([a2d7501](https://github.com/rswanson/lfk/commit/a2d75018e507f2820782a0a458bec678668b3efe))
* remove inert QoS E-core routing ([#7](https://github.com/rswanson/lfk/issues/7)) + energy deep-dive ([#11](https://github.com/rswanson/lfk/issues/11)) ([ee09a4e](https://github.com/rswanson/lfk/commit/ee09a4e2ab916e245ce120392be2238b0801a54f))
* share resource-list cache across hover and drill-in ([1364684](https://github.com/rswanson/lfk/commit/13646845a319e79f4dc774846c5fb019adba392b))
* **ui:** cache layout + rows in middle-column TableRenderer ([c36abe3](https://github.com/rswanson/lfk/commit/c36abe3c0c7dc26aa46453490db9518bcf9fd3d7))
* **ui:** cache layout + rows in middle-column TableRenderer ([12266e4](https://github.com/rswanson/lfk/commit/12266e436b4be004e53874f012be5ec9a9b73e42))


### Reverts

* **app:** remove spinner tick-chain gate from [#206](https://github.com/rswanson/lfk/issues/206) fix ([#215](https://github.com/rswanson/lfk/issues/215)) ([2ab72d2](https://github.com/rswanson/lfk/commit/2ab72d2bd5fcdb7c576d71e6929b910b6abb1eca))


### Code Refactoring

* rename monitoring fallback key from "default" to "_global" ([23a971c](https://github.com/rswanson/lfk/commit/23a971c896e90b82f89dabba4cf175d6275179a4))

## [0.13.2](https://github.com/janosmiko/lfk/compare/v0.13.1...v0.13.2) (2026-05-31)


### Features

* graceful shutdown notice with 10s force-quit timeout ([#314](https://github.com/janosmiko/lfk/issues/314)) ([68e03f6](https://github.com/janosmiko/lfk/commit/68e03f62ce7f8a279112a4f271efbf52228f53f9))

## [0.13.1](https://github.com/janosmiko/lfk/compare/v0.13.0...v0.13.1) (2026-05-31)


### Features

* derive CRD display names from Kind to preserve camel case ([#306](https://github.com/janosmiko/lfk/issues/306)) ([c334c85](https://github.com/janosmiko/lfk/commit/c334c85b571e9d867aaf223bfbc8f27b0ff30cba)), closes [#301](https://github.com/janosmiko/lfk/issues/301)
* per-kind sort memory and per-context column memory ([#310](https://github.com/janosmiko/lfk/issues/310)) ([4e80556](https://github.com/janosmiko/lfk/commit/4e80556ba5e10368ac63804ef5798c80530259bf))


### Bug Fixes

* honor CRD additionalPrinterColumns priority and keep them visible ([#309](https://github.com/janosmiko/lfk/issues/309)) ([bed3a7a](https://github.com/janosmiko/lfk/commit/bed3a7af0d241c34bae675af80d0c1c93544f199))
* link manually-triggered CronJob jobs to their CronJob via ownerRef ([#308](https://github.com/janosmiko/lfk/issues/308)) ([39acb48](https://github.com/janosmiko/lfk/commit/39acb485a3391fc48491481fc09e2a348db52594)), closes [#304](https://github.com/janosmiko/lfk/issues/304)
* remember resource list filter across subview navigation ([#303](https://github.com/janosmiko/lfk/issues/303)) ([#311](https://github.com/janosmiko/lfk/issues/311)) ([4dca616](https://github.com/janosmiko/lfk/commit/4dca616fb12776e7e21ba3f3641e731e82e8e729))

## [0.13.0](https://github.com/janosmiko/lfk/compare/v0.12.9...v0.13.0) (2026-05-30)


### Features

* pin individual resource types into a top "Pinned" section ([#300](https://github.com/janosmiko/lfk/issues/300)) ([3953bd6](https://github.com/janosmiko/lfk/commit/3953bd63416313413705a420b19e8daf32a9b3d2))


### Bug Fixes

* apply theme switch to cached previews immediately ([#299](https://github.com/janosmiko/lfk/issues/299)) ([7d32aaa](https://github.com/janosmiko/lfk/commit/7d32aaa028d3fcc8903d962edb251cdce050076b))
* dashboard events preview background tear under non-black themes ([#293](https://github.com/janosmiko/lfk/issues/293)) ([#296](https://github.com/janosmiko/lfk/issues/296)) ([2d1672d](https://github.com/janosmiko/lfk/commit/2d1672dc7272759cf72904a93982f4c4d4fc644a))

## [0.12.9](https://github.com/janosmiko/lfk/compare/v0.12.8...v0.12.9) (2026-05-29)


### Bug Fixes

* cluster dashboard rendering, scrolling, and responsive usage bars ([#293](https://github.com/janosmiko/lfk/issues/293)) ([#294](https://github.com/janosmiko/lfk/issues/294)) ([9b3fb00](https://github.com/janosmiko/lfk/commit/9b3fb001d4a32f2a7e51780f5dda4e1a1cffae43))

## [0.12.8](https://github.com/janosmiko/lfk/compare/v0.12.7...v0.12.8) (2026-05-28)


### Features

* refresh namespaces in the namespace selector with R ([#292](https://github.com/janosmiko/lfk/issues/292)) ([bc6ace6](https://github.com/janosmiko/lfk/commit/bc6ace6aaf876eda1a5c48260e909a2c421a2db6))


### Bug Fixes

* clear hint-bar message immediately on explorer navigation ([#291](https://github.com/janosmiko/lfk/issues/291)) ([3b5bb3e](https://github.com/janosmiko/lfk/commit/3b5bb3e628cb1aaa55b7de899250372174d707ae))
* gate explorer hint bar by nav level at context picker ([#288](https://github.com/janosmiko/lfk/issues/288)) ([199de4b](https://github.com/janosmiko/lfk/commit/199de4b9730e1a00e349d14b84daec7bb662c874))
* persist in-context read-only toggle to per-context override ([#290](https://github.com/janosmiko/lfk/issues/290)) ([35cbab5](https://github.com/janosmiko/lfk/commit/35cbab5dfdc868303c7300e8a6359a91dc60bb20))

## [0.12.7](https://github.com/janosmiko/lfk/compare/v0.12.6...v0.12.7) (2026-05-28)


### Features

* add negative namespace selection ([#287](https://github.com/janosmiko/lfk/issues/287)) ([653d1cd](https://github.com/janosmiko/lfk/commit/653d1cd9d83708ffa727580af4cd03fc5343c302))


### Bug Fixes

* **nix:** update vendorHash after k8s dependency bump ([#285](https://github.com/janosmiko/lfk/issues/285)) ([9020c87](https://github.com/janosmiko/lfk/commit/9020c87fe35b4d8b60cfae0c23ab8ed9346bef5d)), closes [#284](https://github.com/janosmiko/lfk/issues/284)

## [0.12.6](https://github.com/janosmiko/lfk/compare/v0.12.5...v0.12.6) (2026-05-25)


### Bug Fixes

* **views:** apply GVR-keyed view columns and hide unlisted builtins ([#277](https://github.com/janosmiko/lfk/issues/277)) ([a2df863](https://github.com/janosmiko/lfk/commit/a2df863902a2f58f6c9dd4ed896a210d43da2051)), closes [#262](https://github.com/janosmiko/lfk/issues/262)

## [0.12.5](https://github.com/janosmiko/lfk/compare/v0.12.4...v0.12.5) (2026-05-25)


### Features

* **columns + views:** REV, kubectl-parity audit, k9s-style views config ([#271](https://github.com/janosmiko/lfk/issues/271)) ([837961c](https://github.com/janosmiko/lfk/commit/837961ce44c1de45eca45e17be0a9f916e86dbe2))

## [0.12.4](https://github.com/janosmiko/lfk/compare/v0.12.3...v0.12.4) (2026-05-24)


### Bug Fixes

* **sort:** numeric ordering for percent columns (CPU%, MEM%, */R, */L) ([#273](https://github.com/janosmiko/lfk/issues/273)) ([f50537a](https://github.com/janosmiko/lfk/commit/f50537a8cc91fee15fabad596f671182922a7b58)), closes [#272](https://github.com/janosmiko/lfk/issues/272)

## [0.12.3](https://github.com/janosmiko/lfk/compare/v0.12.2...v0.12.3) (2026-05-23)


### Features

* **actions:** add "Go to Node" to the Pod action menu ([#264](https://github.com/janosmiko/lfk/issues/264)) ([#269](https://github.com/janosmiko/lfk/issues/269)) ([c9190d1](https://github.com/janosmiko/lfk/commit/c9190d1dbc87e5e4f0bb956e058f5be02cdd1b9b))
* **events:** make the events overlay readable by default ([#263](https://github.com/janosmiko/lfk/issues/263)) ([#270](https://github.com/janosmiko/lfk/issues/270)) ([8bce4ea](https://github.com/janosmiko/lfk/commit/8bce4ea2afc704d655e48f9fa7908feaa124bc58))
* **logger:** surface silent failures with dedup to in-app log ([#268](https://github.com/janosmiko/lfk/issues/268)) ([ccfece0](https://github.com/janosmiko/lfk/commit/ccfece0682b69b165c21edae994320d8440bba25))


### Bug Fixes

* **metrics:** fall back to metrics-api when prometheus route fails ([#266](https://github.com/janosmiko/lfk/issues/266)) ([3a79c37](https://github.com/janosmiko/lfk/commit/3a79c377b9abb45019e3029066fe2cee840c4518))
* normalize clipboard line endings to CRLF on Windows ([#261](https://github.com/janosmiko/lfk/issues/261)) ([#265](https://github.com/janosmiko/lfk/issues/265)) ([2b884aa](https://github.com/janosmiko/lfk/commit/2b884aa610a49b90bc3b08ed9ce29919e5d2e2d1))

## [0.12.2](https://github.com/janosmiko/lfk/compare/v0.12.1...v0.12.2) (2026-05-21)


### Bug Fixes

* **bulk:** clear stale bulk-action snapshot on dispatch and cancel ([#257](https://github.com/janosmiko/lfk/issues/257)) ([16a4c60](https://github.com/janosmiko/lfk/commit/16a4c606b632785b84f383251d88b31d12fd7a47))
* stop node metrics column-order flicker (and prevent the whole class) ([#259](https://github.com/janosmiko/lfk/issues/259)) ([7f5c695](https://github.com/janosmiko/lfk/commit/7f5c695020bee63ed35db09ef20992a31cece50a))

## [0.12.1](https://github.com/janosmiko/lfk/compare/v0.12.0...v0.12.1) (2026-05-19)


### Features

* add jump-back navigation history ([#249](https://github.com/janosmiko/lfk/issues/249)) ([#256](https://github.com/janosmiko/lfk/issues/256)) ([c7eae4b](https://github.com/janosmiko/lfk/commit/c7eae4ba6828bcb1abca2a50b63fac6aae449835))


### Bug Fixes

* **port-forward:** reuse local port on restart, fix setup-overlay UX ([#253](https://github.com/janosmiko/lfk/issues/253)) ([d065a5e](https://github.com/janosmiko/lfk/commit/d065a5e15204f9edfb55314a8186420513eb3d59))
* **sort:** sort numeric/structured columns numerically ([#255](https://github.com/janosmiko/lfk/issues/255)) ([1a13ad0](https://github.com/janosmiko/lfk/commit/1a13ad03893d239819238cc4266765a85ef9d8b6)), closes [#250](https://github.com/janosmiko/lfk/issues/250)

## [0.12.0](https://github.com/janosmiko/lfk/compare/v0.11.8...v0.12.0) (2026-05-18)


### Features

* add configurable data directories (LFK_*_DIR overrides) ([#246](https://github.com/janosmiko/lfk/issues/246)) ([4de5317](https://github.com/janosmiko/lfk/commit/4de531735ea2d480165f0d19c0760061ce5ec798))
* add multi-cluster union view with --union-context and --union-set ([#172](https://github.com/janosmiko/lfk/issues/172)) ([ba0f405](https://github.com/janosmiko/lfk/commit/ba0f4059db6cd6588a4e4d8c9c4ba77fadd543db))

## [0.11.8](https://github.com/janosmiko/lfk/compare/v0.11.7...v0.11.8) (2026-05-15)


### Features

* **config:** make kubeconfig discovery directory configurable ([#243](https://github.com/janosmiko/lfk/issues/243)) ([71fddf5](https://github.com/janosmiko/lfk/commit/71fddf5c43e18019c7affd934a72372403c36210))

## [0.11.7](https://github.com/janosmiko/lfk/compare/v0.11.6...v0.11.7) (2026-05-14)


### Bug Fixes

* **logs:** keep cursor at top when older history loads after gg ([#241](https://github.com/janosmiko/lfk/issues/241)) ([add9187](https://github.com/janosmiko/lfk/commit/add918752cdcd067bc4151bbcee28bd4b0bff1c9))

## [0.11.6](https://github.com/janosmiko/lfk/compare/v0.11.5...v0.11.6) (2026-05-14)


### Features

* **copy:** open copy-as picker on Y with YAML / JSON / Table options ([#237](https://github.com/janosmiko/lfk/issues/237)) ([9f0851d](https://github.com/janosmiko/lfk/commit/9f0851d020c96ae40bf50b63b48366ca923fc1b8))
* **ui:** migrate HelmHistory / HelmRollback / DeploymentRollback to OverlayList ([#234](https://github.com/janosmiko/lfk/issues/234)) ([1fbe874](https://github.com/janosmiko/lfk/commit/1fbe87495036b72b7c9ec1646810911086f81450))
* **ui:** unified overlay components — OverlayList, OverlayConfirm, OverlayInput ([#231](https://github.com/janosmiko/lfk/issues/231)) ([5960090](https://github.com/janosmiko/lfk/commit/59600903183f6a7891282e3c23e335567c6b4362))


### Bug Fixes

* **commandbar:** -A flag, pty inline output, exact-match autocomplete ([#235](https://github.com/janosmiko/lfk/issues/235)) ([040a690](https://github.com/janosmiko/lfk/commit/040a6901b17616cb7df3adb0ab98652111eef912))
* **copy:** JSON shortcut chip + partial-success on bulk YAML/JSON copy ([#239](https://github.com/janosmiko/lfk/issues/239)) ([74ec4df](https://github.com/janosmiko/lfk/commit/74ec4df9e173f6a30764015d89ffac765c3f4f75))
* details pane shows children + theme color tracking ([#238](https://github.com/janosmiko/lfk/issues/238)) ([61e0ca3](https://github.com/janosmiko/lfk/commit/61e0ca3f6a42213352d715180cf28cd071ca5e8c))
* **kv-editor:** dedicated selection column for the Secret/ConfigMap/Label editors ([#240](https://github.com/janosmiko/lfk/issues/240)) ([2c00356](https://github.com/janosmiko/lfk/commit/2c00356b486f11e5d50626dbcdf819cb4f84dca2))

## [0.11.5](https://github.com/janosmiko/lfk/compare/v0.11.4...v0.11.5) (2026-05-13)


### Features

* **filters:** add Not Running / Not Bound presets and config invert flag ([#230](https://github.com/janosmiko/lfk/issues/230)) ([3af1652](https://github.com/janosmiko/lfk/commit/3af16528ec28973e4e0893fb47754aa48a050396))


### Bug Fixes

* **nix:** build with Go 1.26.3 by overriding pkgs.go_1_26 ([#228](https://github.com/janosmiko/lfk/issues/228)) ([f4d0ad6](https://github.com/janosmiko/lfk/commit/f4d0ad6c3b8b527af26171fca757886f4ff72d21))

## [0.11.4](https://github.com/janosmiko/lfk/compare/v0.11.3...v0.11.4) (2026-05-12)


### Features

* **karpenter:** first-class actions for NodePool / NodeClaim / EC2NodeClass ([#223](https://github.com/janosmiko/lfk/issues/223)) ([5f37b70](https://github.com/janosmiko/lfk/commit/5f37b70e21a4edf09aee9416332784e41ed3c15a))
* **knative:** first-class Knative Serving (Activate) + Eventing icons ([#224](https://github.com/janosmiko/lfk/issues/224)) ([e89be5b](https://github.com/janosmiko/lfk/commit/e89be5baa1405e366334a52bb999e494aad63e74))


### Bug Fixes

* **app:** respect KUBE_EDITOR and parse editor flags ([#226](https://github.com/janosmiko/lfk/issues/226)) ([d944265](https://github.com/janosmiko/lfk/commit/d9442656a5be06685742a63596bfd867ba879fb4))
* **scheduler:** stop title-bar spinner during 10s linger window after work completes ([#220](https://github.com/janosmiko/lfk/issues/220)) ([f441c78](https://github.com/janosmiko/lfk/commit/f441c7844e3bcd0450a3c1f11445633ed7879c79))

## [0.11.3](https://github.com/janosmiko/lfk/compare/v0.11.2...v0.11.3) (2026-05-12)


### Bug Fixes

* **app:** refresh right-pane preview at LevelResourceTypes on tab switch and watch tick ([#216](https://github.com/janosmiko/lfk/issues/216)) ([cc4c90f](https://github.com/janosmiko/lfk/commit/cc4c90f68d5e4d73c6e8915bf340fd61335c73b8))
* **ui:** drop blank line between RESOURCE USAGE header and bars ([#217](https://github.com/janosmiko/lfk/issues/217)) ([40be2bc](https://github.com/janosmiko/lfk/commit/40be2bc9148c41d961e1fadd0c786dfb8854a150))


### Reverts

* **app:** remove spinner tick-chain gate from [#206](https://github.com/janosmiko/lfk/issues/206) fix ([#215](https://github.com/janosmiko/lfk/issues/215)) ([2ab72d2](https://github.com/janosmiko/lfk/commit/2ab72d2bd5fcdb7c576d71e6929b910b6abb1eca))

## [0.11.2](https://github.com/janosmiko/lfk/compare/v0.11.1...v0.11.2) (2026-05-11)


### Bug Fixes

* **app,scheduler:** drop idle CPU from ~145% to ~0% (closes [#206](https://github.com/janosmiko/lfk/issues/206)) ([#211](https://github.com/janosmiko/lfk/issues/211)) ([bcdd3d0](https://github.com/janosmiko/lfk/commit/bcdd3d0ce10cb924a4e58703cc64367ddb7c5688))

## [0.11.1](https://github.com/janosmiko/lfk/compare/v0.11.0...v0.11.1) (2026-05-11)


### Features

* **clipboard:** support Windows and Wayland via atotto/clipboard ([#195](https://github.com/janosmiko/lfk/issues/195)) ([c1871de](https://github.com/janosmiko/lfk/commit/c1871de47df2a9597f13c7421980095b5d2d8b2c))


### Bug Fixes

* **describe:** route keys to search input, not global tab handler ([#203](https://github.com/janosmiko/lfk/issues/203)) ([#204](https://github.com/janosmiko/lfk/issues/204)) ([40ea18d](https://github.com/janosmiko/lfk/commit/40ea18de1aeb7ad4dc3c5ee29573da9688d7834d))
* **exec,browser:** make interactive shell + browser-open actions work on Windows ([#197](https://github.com/janosmiko/lfk/issues/197)) ([4954439](https://github.com/janosmiko/lfk/commit/4954439941e12860973643a0f6a91b36232a713e))
* **release:** skip Chocolatey publish until first version is moderated ([#201](https://github.com/janosmiko/lfk/issues/201)) ([0a5be5d](https://github.com/janosmiko/lfk/commit/0a5be5d8d3fa0082576a2d04f68bf42538c6b7cc))

## [0.11.0](https://github.com/janosmiko/lfk/compare/v0.10.4...v0.11.0) (2026-05-09)


### Features

* **k8s:** surface ephemeral containers in pod views ([#180](https://github.com/janosmiko/lfk/issues/180)) ([ac1a1c5](https://github.com/janosmiko/lfk/commit/ac1a1c54baf293482a7a29666336816838713332))
* **localcluster:** manage kind/k3d/minikube clusters from inside lfk ([#175](https://github.com/janosmiko/lfk/issues/175)) ([3c85fd9](https://github.com/janosmiko/lfk/commit/3c85fd9955f5fd88dc78efefd298cac11f0e6bf3))
* **release:** add AUR channel (lfk-bin) ([#174](https://github.com/janosmiko/lfk/issues/174)) ([c6df49d](https://github.com/janosmiko/lfk/commit/c6df49ddd153d0d15148e0dfa181f9440a19a702))
* **scheduler:** priority task queue with per-context dispatch ([#186](https://github.com/janosmiko/lfk/issues/186)) ([80e0ba1](https://github.com/janosmiko/lfk/commit/80e0ba13f5dc7d4a3609e3233002752c6686d8c6))
* traffic capture (kubectl-debug + kubeshark backends) ([#179](https://github.com/janosmiko/lfk/issues/179)) ([b51d64c](https://github.com/janosmiko/lfk/commit/b51d64c421a4001ff1cf1c53efc2792cceb41b3f))
* **viewers:** vim text-object selection (viw/vaw/viW/vaW) ([#185](https://github.com/janosmiko/lfk/issues/185)) ([7eb0aea](https://github.com/janosmiko/lfk/commit/7eb0aeaf2b5263e7c7fd630b3ac300b7e5140061))


### Bug Fixes

* **actions:** block delete keypress in containers view ([#181](https://github.com/janosmiko/lfk/issues/181)) ([584ff4f](https://github.com/janosmiko/lfk/commit/584ff4fd0bf50f3ba1c65b061469bcf741e09804))
* **nodeshell:** land on DiskPressure/MemoryPressure/PIDPressure nodes ([#177](https://github.com/janosmiko/lfk/issues/177)) ([eec8d02](https://github.com/janosmiko/lfk/commit/eec8d02a5426c1667f9d0a195bfa8765da1c98d2))
* stop infinite Loading spinner on permission errors ([#171](https://github.com/janosmiko/lfk/issues/171)) ([07c4c14](https://github.com/janosmiko/lfk/commit/07c4c14c90800144ad5999218fbfdf887069dfc9))
* **tabs:** refresh middle column on tab switch (stale-while-revalidate) ([#182](https://github.com/janosmiko/lfk/issues/182)) ([b84595e](https://github.com/janosmiko/lfk/commit/b84595e5718c663ed9b6096f4e0f6e3960f065d0))
* **ui:** clip pinned resource-usage footer no longer triggered by event count ([#178](https://github.com/janosmiko/lfk/issues/178)) ([54b6af0](https://github.com/janosmiko/lfk/commit/54b6af0cdf2c56192ec15bc6233c4489cdf3753f))

## [0.10.4](https://github.com/janosmiko/lfk/compare/v0.10.3...v0.10.4) (2026-05-06)


### Features

* **argocd:** add Sync Wave Timeline overlay ([#160](https://github.com/janosmiko/lfk/issues/160)) ([3784fc6](https://github.com/janosmiko/lfk/commit/3784fc6e3de25fa3774457f487b0c5840e01131f))
* **release:** add cloudsmith deb+rpm channel ([#163](https://github.com/janosmiko/lfk/issues/163)) ([8c50bec](https://github.com/janosmiko/lfk/commit/8c50bec5d2e7fb158d960543939b57b171b2915a))
* **release:** add scoop, winget, chocolatey channels ([#161](https://github.com/janosmiko/lfk/issues/161)) ([f4fe4a5](https://github.com/janosmiko/lfk/commit/f4fe4a576486b3c84e1949f773e6cfc3bad0ad84))

## [0.10.3](https://github.com/janosmiko/lfk/compare/v0.10.2...v0.10.3) (2026-05-06)


### Features

* **release:** foundation for new package-manager channels ([#159](https://github.com/janosmiko/lfk/issues/159)) ([0a4a353](https://github.com/janosmiko/lfk/commit/0a4a353a3e715b05a6b5b86e2a4183e26c9b21e9))
* **ui:** dim explorer behind overlays via dim_overlay option ([#99](https://github.com/janosmiko/lfk/issues/99)) ([df167f4](https://github.com/janosmiko/lfk/commit/df167f4bb20f65eda38b72ff80424755d8deb8d5))


### Bug Fixes

* **filter:** clear active filter preset on Esc ([#156](https://github.com/janosmiko/lfk/issues/156)) ([7b22dff](https://github.com/janosmiko/lfk/commit/7b22dff6d59d31d8484536abe445cd6774be45a8))
* **filter:** clear stale preview when filter preset matches zero items ([#157](https://github.com/janosmiko/lfk/issues/157)) ([5a8f28c](https://github.com/janosmiko/lfk/commit/5a8f28c79aa2579ab184a42b8184855f059469f9))
* **theme:** keep parent highlight readable on themes with near-text border ([b5fc86f](https://github.com/janosmiko/lfk/commit/b5fc86f4ea5b9abb0c1f54a10d356395a8f907a2))

## [0.10.2](https://github.com/janosmiko/lfk/compare/v0.10.1...v0.10.2) (2026-05-05)


### Bug Fixes

* **release:** rename cosign bundle to .sigstore for Scorecard ([#152](https://github.com/janosmiko/lfk/issues/152)) ([f12ce39](https://github.com/janosmiko/lfk/commit/f12ce3976ceae1fd6255d573de842647a687a1a4))

## [0.10.1](https://github.com/janosmiko/lfk/compare/v0.10.0...v0.10.1) (2026-05-05)


### Bug Fixes

* **release:** declare cosign bundle as signature artifact ([#150](https://github.com/janosmiko/lfk/issues/150)) ([1adf6ea](https://github.com/janosmiko/lfk/commit/1adf6eae58e26f23510291c0727e60cc8b24da60))

## [0.10.0](https://github.com/janosmiko/lfk/compare/v0.9.39...v0.10.0) (2026-05-05)


### ⚠ BREAKING CHANGES

* add multi-strategy right-sizing advisor overlay ([#148](https://github.com/janosmiko/lfk/issues/148))
* CrashLoopBackOff investigator overlay

### Features

* add multi-strategy right-sizing advisor overlay ([#148](https://github.com/janosmiko/lfk/issues/148)) ([5392610](https://github.com/janosmiko/lfk/commit/539261090646f1dc94c19dcd3c1b57eca1e7b1bb))
* CrashLoopBackOff investigator overlay ([93d310e](https://github.com/janosmiko/lfk/commit/93d310e8eb2b6547c8967749aec7a5e5a318f9ef))


### Bug Fixes

* **metrics:** stop ~1Hz column-order blink on PodInitializing rows ([0895f56](https://github.com/janosmiko/lfk/commit/0895f563667f9ae1cd25a08d2f0ac33a5a21f111))
* **metrics:** stop ~1Hz column-order blink on PodInitializing rows ([b1b53cf](https://github.com/janosmiko/lfk/commit/b1b53cf05b1bc8fbddd0188b80d9727f162d3f76))
* **release:** migrate cosign signing to Sigstore bundle output ([acc6284](https://github.com/janosmiko/lfk/commit/acc62843638f64d1ca38a762cd1cc05295bf43e2))
* **ui:** stop namespace and column-toggle overlays from shrinking on filter ([7965ecd](https://github.com/janosmiko/lfk/commit/7965ecd49b03690512fe2d3622a1062a3a74cec1))
* **ui:** stop selector overlays from shrinking on filter ([39d0ba2](https://github.com/janosmiko/lfk/commit/39d0ba2a5a5ff952dfc2dd6bcbc170e78ced4dfa))

## [0.9.39](https://github.com/janosmiko/lfk/compare/v0.9.38...v0.9.39) (2026-05-04)


### Features

* **clusters:** add per-cluster color coding with title-bar tint ([#124](https://github.com/janosmiko/lfk/issues/124)) ([65da3ac](https://github.com/janosmiko/lfk/commit/65da3ac010f4b84b4270dc4ce8662243a7171497))
* **editors:** revamp edit pane — bordered fields + non-shifting cursor ([55b322a](https://github.com/janosmiko/lfk/commit/55b322a33113fbff87ec296a51100e8b5fb41a4e))
* **editors:** wire `s` multi-select + Shift+Y format-copy on ConfigMap + Label editors ([44429de](https://github.com/janosmiko/lfk/commit/44429def06b4106787b14dcb5fc15d907c8741be))
* **editors:** wire `s` multi-select + Shift+Y format-copy on Secret editor ([f240d35](https://github.com/janosmiko/lfk/commit/f240d358e0417e787eaac477b9b3162d42a4b71d))
* **mouse:** click-to-drill, right-click action menu, overlay mouse ([8287ba0](https://github.com/janosmiko/lfk/commit/8287ba0b3fa5c50f462f063293c5915f3871a51c))
* **networking:** per-endpoint preview for Endpoints / EndpointSlices ([fb0201d](https://github.com/janosmiko/lfk/commit/fb0201d9513a19ff9c3ad429a0a331b5916f9af5))
* **networking:** Service preview rollup of backing EndpointSlices ([66a1e26](https://github.com/janosmiko/lfk/commit/66a1e26b74894eafd53b746012bd8eeb50586c40))
* **rbac:** reverse-RBAC "Who-Can" view, layered on the Can-I overlay ([7598c68](https://github.com/janosmiko/lfk/commit/7598c68501a0db5067614c8cbf35f3e91020d7fc))
* **resource-map:** traverse Pod refs with MissingRef detection ([a0517c3](https://github.com/janosmiko/lfk/commit/a0517c306964e02bfd8a3126688727ec4f96bf77))
* **viewers:** extend count-prefix to column / word / page / search motions ([64e9498](https://github.com/janosmiko/lfk/commit/64e9498894acd84684a57679b9387c364de36f56))
* **viewers:** match vim/nvim [count]&lt;C-d&gt;/&lt;C-u&gt; 'scroll' option semantics ([fd83a57](https://github.com/janosmiko/lfk/commit/fd83a57cd57c5bc950d507ad78ce7c38f0e00bec))


### Bug Fixes

* address CodeRabbit findings on PR [#122](https://github.com/janosmiko/lfk/issues/122) ([1a0a97c](https://github.com/janosmiko/lfk/commit/1a0a97cedafa22d2f623b1cb5091c73f2fd6d462))
* address second round of CodeRabbit findings on PR [#122](https://github.com/janosmiko/lfk/issues/122) ([3fffdd4](https://github.com/janosmiko/lfk/commit/3fffdd4ff336da514768307b131c79067e696252))
* clear stale pod metrics when metrics-server payload is empty ([0191775](https://github.com/janosmiko/lfk/commit/01917750e54e90baba1e1694c4dd301f8f39d50a))
* **editors:** address CodeRabbit review on PR [#134](https://github.com/janosmiko/lfk/issues/134) ([c18c1ce](https://github.com/janosmiko/lfk/commit/c18c1ce1debc0296747f146282803c7d9af3ff88))
* **editors:** ANSI leak in field labels + up/down nav + scroll-to-cursor ([b82144c](https://github.com/janosmiko/lfk/commit/b82144c9a22b0e647ba984d9aa66bcec76f440ce))
* **editors:** collapse long/multi-line values to a single visual cell ([97481d5](https://github.com/janosmiko/lfk/commit/97481d5aa6ebf98798a2e12caba335a081f409f9))
* **editors:** consistent key column + space-select + smart-y ([1302e3b](https://github.com/janosmiko/lfk/commit/1302e3b46cb8e31149f109080298faba9adb9c13))
* **editors:** ctrl+s under active filter no longer mutates wrong key ([c02de2b](https://github.com/janosmiko/lfk/commit/c02de2b35d6bf09ee4a538aa1bc5726ebbd01663))
* **editors:** cursor in edit pane lands at TextInput cursor pos + ([6ab3526](https://github.com/janosmiko/lfk/commit/6ab35266d90fa34f6a1d82b53d32ee91269d5b5c))
* **editors:** format picker no longer shrinks the table ([1e0a2eb](https://github.com/janosmiko/lfk/commit/1e0a2eb1d087a0e8b25fe39b5d48b501c498a50d))
* **editors:** inline edit mode for single-line values ([2cc3e53](https://github.com/janosmiko/lfk/commit/2cc3e53be4b80e377ff9903510d68b35b63d368b))
* **editors:** show multi-line values as multi-line during editing ([6c59ceb](https://github.com/janosmiko/lfk/commit/6c59ceb9c4d79ef96d72655fd250e5fca10b1e19))
* **editors:** sticky scroll + ctrl+u/d/f/b page keys + line-scoped ctrl+a/e ([1d8ba02](https://github.com/janosmiko/lfk/commit/1d8ba021ab7ba45db07bbc1a4ea91235b2f0e313))
* **help:** address lint and CodeRabbit review ([490fe6d](https://github.com/janosmiko/lfk/commit/490fe6daf363b8011508a08fea8f2ab321eaf70d))
* **mouse:** address CodeRabbit review on PR [#135](https://github.com/janosmiko/lfk/issues/135) ([ed9293b](https://github.com/janosmiko/lfk/commit/ed9293b2e3c9741e8afff240be802d253db74ad9))
* **nav:** preserve cursor on watch-tick discovery failure at LevelResourceTypes ([057f036](https://github.com/janosmiko/lfk/commit/057f036fb904e2df8494ebdf65b1cb65796745b1))
* **networking:** address coderabbit findings on Service endpoints rollup ([de75cb3](https://github.com/janosmiko/lfk/commit/de75cb31dfdfa87d1134ddc05ab9b64064680293))
* **networking:** always refetch Service endpoints; cache hid pod churn ([906b37f](https://github.com/janosmiko/lfk/commit/906b37fb80b547be611f75203e8546352acfea4e))
* **networking:** carry over Service rollup columns across watch-tick rebuilds ([e67c272](https://github.com/janosmiko/lfk/commit/e67c272da3f17da5b16b35bb2606dea8600882ed))
* **networking:** stale-while-revalidate Service endpoints to stop the flash ([24b6f95](https://github.com/janosmiko/lfk/commit/24b6f95e5b78806787693c3150883b57168f3472))
* **networking:** treat absent EndpointSlice conditions.ready as ready ([94ac9fc](https://github.com/janosmiko/lfk/commit/94ac9fc08c66a5691ff68df90b33691a49b926e9))
* **preview:** clear previewLoading when resource list arrives empty ([a781377](https://github.com/janosmiko/lfk/commit/a781377429081c8794ce7ccd20674b7f03cf031a))
* **preview:** DATA (N) counts keys not visual lines ([fa623d4](https://github.com/janosmiko/lfk/commit/fa623d4bbff9a39749a90060f2bb4d577faf6934))
* **quit:** cancel in-flight API requests so quit doesn't hang on dead clusters ([b0479b8](https://github.com/janosmiko/lfk/commit/b0479b8978a32387c15452cf0ff077e9e005664c))
* **rbac:** address remaining coderabbit findings on Who-Can ([8acdd11](https://github.com/janosmiko/lfk/commit/8acdd11965492c7f1d1f6a328ad7aa1e514678b4))
* **resource-map:** fall back to nav.Namespace at LevelContainers ([13033f3](https://github.com/janosmiko/lfk/commit/13033f3195f66bb5c521f6c2aa532b467681dcad))
* **resource-map:** show Pod's tree when M is pressed at LevelContainers ([36b84fa](https://github.com/janosmiko/lfk/commit/36b84fa7f81ee5dd57e386d407997f7d58f349a8))
* **tabs:** persist right-pane footers per tab so metrics don't bleed ([1a4fa9d](https://github.com/janosmiko/lfk/commit/1a4fa9d20f3b225e3e727dd0dd72269928ad1b0f))
* **viewers:** clear diff digit buffer on visual mode entry ([bac6309](https://github.com/janosmiko/lfk/commit/bac63095c5ae74baff9dfd0a409362ae63ae48ec))
* **viewers:** round half-page step before scaling by count ([316a91d](https://github.com/janosmiko/lfk/commit/316a91dffd2446afde5c21d69f1e9b0147682c38))
* **viewers:** scale yaml page motions by viewport, not raw m.height ([67dd610](https://github.com/janosmiko/lfk/commit/67dd6102ad5473f136723107fc1eed9690e4206e))

## [0.9.38](https://github.com/janosmiko/lfk/compare/v0.9.37...v0.9.38) (2026-05-02)


### Features

* **logs:** persistent search history with Up/Down recall ([58d6b08](https://github.com/janosmiko/lfk/commit/58d6b08693996fb7d292b471de5159d424133119))
* **logs:** persistent search history with Up/Down recall in log viewer ([cc70537](https://github.com/janosmiko/lfk/commit/cc7053710e8f338aa7bbc70b66ae32d8f4f5c5d6))


### Bug Fixes

* **history:** preserve draft on edit-after-recall via leaveBrowse() ([321a4bc](https://github.com/janosmiko/lfk/commit/321a4bc4e75c62bc3a3110996fe8a3f687600352))
* **history:** tighten file perms and leaveBrowse on paste ([0678df4](https://github.com/janosmiko/lfk/commit/0678df45075c3b136359a42db236ff1ee2f673e7))
* **logs:** handle Ctrl+U (delete-line) in log viewer search input ([cd1a049](https://github.com/janosmiko/lfk/commit/cd1a0495721989cb574777ccd6876242870db6f0))
* **logs:** scope log-search backspace reset() inside len-guard ([6608c14](https://github.com/janosmiko/lfk/commit/6608c14640902b22495c82a03709c2410978cab4))

## [0.9.37](https://github.com/janosmiko/lfk/compare/v0.9.36...v0.9.37) (2026-05-02)


### Features

* **ui:** support count-prefixed motion (Nj/Nk) in read-only viewers ([1068839](https://github.com/janosmiko/lfk/commit/10688392fb6bf4b49d6a8b20bd3ada5ad1b3335f))
* **ui:** support count-prefixed yank (Ny) in read-only viewers ([28782d1](https://github.com/janosmiko/lfk/commit/28782d1e9064fe83339192d8f406e7d93b516c9a))


### Bug Fixes

* **nav:** clear filter state when navigating to parent ([9977274](https://github.com/janosmiko/lfk/commit/99772748fba84eb64889d63bd417dd0f597cf007))
* **ui:** keep "/" search highlight from corrupting SGR codes ([8383f6b](https://github.com/janosmiko/lfk/commit/8383f6bed5fb9840f331b3510af117911c5dc4b9))

## [0.9.36](https://github.com/janosmiko/lfk/compare/v0.9.35...v0.9.36) (2026-04-30)


### Features

* **k8s:** cache resource lists via shared informer (closes [#86](https://github.com/janosmiko/lfk/issues/86)) ([c8578cc](https://github.com/janosmiko/lfk/commit/c8578cc3fb4b7997a780741ddbe14e54b6e807fe))
* **ui:** advertise y/n alongside Enter/Esc for confirm dialogs ([4f95b5c](https://github.com/janosmiko/lfk/commit/4f95b5cf4140371c1cfa00fb4cf78b1faca9c779))
* **ui:** pin info chips far-right + entry-aware keymap fit ([#101](https://github.com/janosmiko/lfk/issues/101)) ([1163c7a](https://github.com/janosmiko/lfk/commit/1163c7a9ea928e484d75b500f7db3029035ac029))


### Bug Fixes

* **app:** keep silent ns refresh from clobbering an open overlay ([73d5ba2](https://github.com/janosmiko/lfk/commit/73d5ba2dd8483ddbc2980559ca1a06306097ea5f))
* **app:** make node shell work on SELinux-enforcing immutable distros ([4477e9c](https://github.com/janosmiko/lfk/commit/4477e9cc516d269dc23093c9af8c957a9c2a1689))
* **k8s,ui:** harden informer cache wiring + config parsing ([ddd596e](https://github.com/janosmiko/lfk/commit/ddd596eb2f21f86d5d863aab50a4dc9c803ac60b))
* **ui:** center quit overlay text and unify confirm-hint convention ([e200ffe](https://github.com/janosmiko/lfk/commit/e200ffe6cb0d140a749d695de86420dca72e5c79))
* **ui:** drop dangling CONTRIBUTING.md refs and tighten confirm-hint test ([315e04d](https://github.com/janosmiko/lfk/commit/315e04d07c0dbca699987f4db6ad042dfa064e59))
* **ui:** invalidate middle-column row cache on theme change ([01ca28f](https://github.com/janosmiko/lfk/commit/01ca28f6963e6c0cffc9203c79a44916d67532a4))
* **ui:** widen Quick Filters overlay and clean up selected row ([86afe1a](https://github.com/janosmiko/lfk/commit/86afe1acb04a3e2bb7e0fee670c931e999211a7c))


### Performance Improvements

* **app:** seed namespace selector overlay from existing cache ([d8090b4](https://github.com/janosmiko/lfk/commit/d8090b42ab45cf6f485f84957d8907cd6f136a75))

## [0.9.35](https://github.com/janosmiko/lfk/compare/v0.9.34...v0.9.35) (2026-04-29)


### Features

* **app:** tackle PTY pain points from [#81](https://github.com/janosmiko/lfk/issues/81) — selection, mux mode, scrollback ([32be754](https://github.com/janosmiko/lfk/commit/32be7546df5bf3df871e7a9d9d38a45e912452b1))


### Bug Fixes

* **app:** require typed confirmation for action-menu Force Delete ([8b1b2b7](https://github.com/janosmiko/lfk/commit/8b1b2b768f37fc23d3bf0fab1143fe11b7e15a81)), closes [#89](https://github.com/janosmiko/lfk/issues/89)
* **app:** unify Force Delete help text across menus, dialogs, and docs ([41b3aba](https://github.com/janosmiko/lfk/commit/41b3abac191aa8ad7cc50b1efe37e83257815219))

## [0.9.34](https://github.com/janosmiko/lfk/compare/v0.9.33...v0.9.34) (2026-04-29)


### Features

* **app:** add read-only mode with per-context [RO] markers ([1b1d9c1](https://github.com/janosmiko/lfk/commit/1b1d9c1738db93ea1b82f9979e8eaef51764832a))
* **app:** add read-only mode with per-context [RO] markers ([c148097](https://github.com/janosmiko/lfk/commit/c148097833a0771762632562b5b9066c696f7f6d))
* **app:** apply y/Y to multi-selection ([ce71b97](https://github.com/janosmiko/lfk/commit/ce71b97e6acba73d8ffc4cd35a1f669212735925))
* **app:** route :export through the Y bulk dispatcher ([d550328](https://github.com/janosmiko/lfk/commit/d55032814bfed5d718126e1624bc59ee6929487d))
* **ui:** add y to copy cursor row from rollback / history overlays ([5cc7cf6](https://github.com/janosmiko/lfk/commit/5cc7cf64009ac20f9e826d2373274479c5f717d9))


### Bug Fixes

* **app:** apply Y bulk to LevelOwned and skip false bulk at LevelContainers ([e366579](https://github.com/janosmiko/lfk/commit/e366579924f2d286702f7d6d0c9bf49f2826e949))
* **app:** plug read-only bypasses across labels and overlays ([66a11d0](https://github.com/janosmiko/lfk/commit/66a11d0f421d9cde78b874e04402299b373df3f2))
* **ui:** gate :sort command and column-header clicks on sortApplies() ([f8a7941](https://github.com/janosmiko/lfk/commit/f8a794135bd5f6f8a77ae7a109aad6b862e3a9db))
* **ui:** hide no-op sort and actions at picker levels ([7376449](https://github.com/janosmiko/lfk/commit/7376449eb354be3449f3080a1ef1f65120f3abb6))

## [0.9.33](https://github.com/janosmiko/lfk/compare/v0.9.32...v0.9.33) (2026-04-28)


### Bug Fixes

* **ui:** show full hotkey hint bar with log preview on ([#71](https://github.com/janosmiko/lfk/issues/71)) ([0badd03](https://github.com/janosmiko/lfk/commit/0badd0302d56738c7c6934fc38b3c4f457ac4e83))


### Performance Improvements

* **discovery:** persist API discovery to disk for stale-while-revalidate startup ([a1aaf27](https://github.com/janosmiko/lfk/commit/a1aaf27622f1809ab7bef9214f3cff19e3235399))
