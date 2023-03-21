# Changelog

## [0.2.32](https://github.com/james-milligan/open-feature-operator/compare/v0.2.31...v0.2.32) (2023-03-21)


### 🔄 Refactoring

* logging cleanup ([#308](https://github.com/james-milligan/open-feature-operator/issues/308)) ([75bdd8b](https://github.com/james-milligan/open-feature-operator/commit/75bdd8ba6292e432e0ed422398c69387d44f8ce6))
* update file mount path to conform with schema store regex ([#331](https://github.com/james-milligan/open-feature-operator/issues/331)) ([56ea082](https://github.com/james-milligan/open-feature-operator/commit/56ea082c99cf1f4ef1552ac79e3f5f5d2a883221))


### ✨ New Features

* add ff shortname, commit httpSyncConfiguration ([11e4652](https://github.com/james-milligan/open-feature-operator/commit/11e46528fcd06cdc0c8e6f46944656224cd97441))
* add log format configuration options through helm chart ([#346](https://github.com/james-milligan/open-feature-operator/issues/346)) ([bcef736](https://github.com/james-milligan/open-feature-operator/commit/bcef7368fc4905b351f81f5dfa10eb4c26bf8764))
* adding artifacthub information ([#144](https://github.com/james-milligan/open-feature-operator/issues/144)) ([65a5244](https://github.com/james-milligan/open-feature-operator/commit/65a524445d1db8bb5608b88282a4d97a9bb6b74f))
* backfill flagd-kubernetes-sync cluster role binding on startup ([#299](https://github.com/james-milligan/open-feature-operator/issues/299)) ([02ca541](https://github.com/james-milligan/open-feature-operator/commit/02ca541353afdbfe4cdc41f32f9399826d4f6fc5))
* builds helm chart ([#137](https://github.com/james-milligan/open-feature-operator/issues/137)) ([1525421](https://github.com/james-milligan/open-feature-operator/commit/1525421229d43b17636dddb65d7b124e6477fe79))
* default sync provider configuration ([#320](https://github.com/james-milligan/open-feature-operator/issues/320)) ([7cba7e1](https://github.com/james-milligan/open-feature-operator/commit/7cba7e14c223a083f02ff8313b899583253120f3))
* enable flagd probes ([#390](https://github.com/james-milligan/open-feature-operator/issues/390)) ([41efb15](https://github.com/james-milligan/open-feature-operator/commit/41efb155994b3cfb768cc39e59bfc09781c57f2e))
* Helm configuration ([#304](https://github.com/james-milligan/open-feature-operator/issues/304)) ([99edfeb](https://github.com/james-milligan/open-feature-operator/commit/99edfeb8c32ada435f830c6799540ebdf3b5fcdd))
* improve deployment pattern ([#344](https://github.com/james-milligan/open-feature-operator/issues/344)) ([572ba96](https://github.com/james-milligan/open-feature-operator/commit/572ba961912ada2c07eb6143925d16ab6a6a85a3))
* introduce configurable resource limits for flagd sidecar ([e4affcf](https://github.com/james-milligan/open-feature-operator/commit/e4affcfb0ccf13dc0406ef1c21c2b884a836f71f))
* Introduced context to the readyz endpoint, added wait to test suite ([#336](https://github.com/james-milligan/open-feature-operator/issues/336)) ([ed81c02](https://github.com/james-milligan/open-feature-operator/commit/ed81c0284f8d759eb228d3af7030efb0b94ee280))
* introduced v1beta1 of featureflagconfiguration CRD with conversion webhook to v1alpha1 ([a45bdef](https://github.com/james-milligan/open-feature-operator/commit/a45bdef5eec87738ce731af5825daffeb69eb6cb))
* metrics ([#111](https://github.com/james-milligan/open-feature-operator/issues/111)) ([6016669](https://github.com/james-milligan/open-feature-operator/commit/6016669ec46984d127951ee5d0ff02e7685f4d80))
* pr github action workflow ([#96](https://github.com/james-milligan/open-feature-operator/issues/96)) ([a719f8a](https://github.com/james-milligan/open-feature-operator/commit/a719f8a33abc9b9599987314282cc4e7ac202d67))
* stop creation and mounting of flagd-config config map in case of kubernetes sync-provider ([#126](https://github.com/james-milligan/open-feature-operator/issues/126)) ([a1d9fe2](https://github.com/james-milligan/open-feature-operator/commit/a1d9fe276a37259d01e6ed6239c0ebcd3a1e6611))
* structured the featureflagconfiguration CRD ([b056c7c](https://github.com/james-milligan/open-feature-operator/commit/b056c7cdd76f4653c1a728342687beaa8279e314))
* upgrade flagd to v0.3.0 ([20571e1](https://github.com/james-milligan/open-feature-operator/commit/20571e1018e102ffbcf01b2518fcbf8b66a287be))


### 📚 Documentation

* add AND operator to sequential commands ([#368](https://github.com/james-milligan/open-feature-operator/issues/368)) ([6f73a62](https://github.com/james-milligan/open-feature-operator/commit/6f73a6214d87771f9555469fe4d60dbb6d301198))
* automated demo deployment ([#310](https://github.com/james-milligan/open-feature-operator/issues/310)) ([e9cb659](https://github.com/james-milligan/open-feature-operator/commit/e9cb6598fadc440ed300231dbacf88926945724a))
* fix link to artifact hub ([9e9590b](https://github.com/james-milligan/open-feature-operator/commit/9e9590b1d60b87b96a90f69ae89401362cc23d9b))
* fix rendering issue with operator resource config table ([#401](https://github.com/james-milligan/open-feature-operator/issues/401)) ([71ea8a6](https://github.com/james-milligan/open-feature-operator/commit/71ea8a68bbb97052822552ffce3c498c3da0e52d))
* improve getting started guide ([#288](https://github.com/james-milligan/open-feature-operator/issues/288)) ([a5a009b](https://github.com/james-milligan/open-feature-operator/commit/a5a009b8ea10dabf4c2034c54b1eb1e8d71e7c23))
* Update permissions docs ([#350](https://github.com/james-milligan/open-feature-operator/issues/350)) ([e89ea75](https://github.com/james-milligan/open-feature-operator/commit/e89ea7528d1e7546b2b4c0a81a105dda309f52b3))


### 🐛 Bug Fixes

* add sbom to ouput name ([#182](https://github.com/james-milligan/open-feature-operator/issues/182)) ([5e939a8](https://github.com/james-milligan/open-feature-operator/commit/5e939a8f67fbd095c18a6a2172bb856fe61dd173))
* artifact name and output file ([#187](https://github.com/james-milligan/open-feature-operator/issues/187)) ([4dee157](https://github.com/james-milligan/open-feature-operator/commit/4dee157d44c20fc925f9e33dbaae16c18f3d9b48))
* build and push to docker registry on tag creation ([#121](https://github.com/james-milligan/open-feature-operator/issues/121)) ([27c6f9c](https://github.com/james-milligan/open-feature-operator/commit/27c6f9cbc298fb8bf578464e4c3f9f07402b87ab))
* build and push to docker registry with tag as current release ([#123](https://github.com/james-milligan/open-feature-operator/issues/123)) ([d4abda1](https://github.com/james-milligan/open-feature-operator/commit/d4abda119e4a7c2dab7a2e0d335d44b1df07ec62))
* bump flagd version to include change detection fix ([421cab6](https://github.com/james-milligan/open-feature-operator/commit/421cab651f6ebe2ece1380fda7dc24d92838d6b5))
* correcrt needs in asset release ([5ed4571](https://github.com/james-milligan/open-feature-operator/commit/5ed45718ca189a15f7cdf4f8ddfc5864f189b1ce))
* CVE-2022-32149 ([015c19a](https://github.com/james-milligan/open-feature-operator/commit/015c19ac4455673902c365111816b021f893c485))
* **deps:** update github.com/open-feature/schemas digest to 302d0fa ([#246](https://github.com/james-milligan/open-feature-operator/issues/246)) ([7d22374](https://github.com/james-milligan/open-feature-operator/commit/7d22374afb7a5e2e166550544d327ec7b5b3d1bf))
* **deps:** update github.com/open-feature/schemas digest to d638ecf ([a984836](https://github.com/james-milligan/open-feature-operator/commit/a98483696f467270783858046132f02b3d338ac2))
* **deps:** update kubernetes packages to v0.25.3 ([8d6103d](https://github.com/james-milligan/open-feature-operator/commit/8d6103d8d41e9fc6f3854ef6a27293a754160ff5))
* **deps:** update kubernetes packages to v0.25.4 ([75bab2d](https://github.com/james-milligan/open-feature-operator/commit/75bab2d441c945d51f17f0d32195a217072c3c15))
* **deps:** update module github.com/open-feature/schemas to v0.2.8 ([#269](https://github.com/james-milligan/open-feature-operator/issues/269)) ([ed48060](https://github.com/james-milligan/open-feature-operator/commit/ed48060b1f9e591ddadca4f9478728a823e10685))
* **deps:** update module go.uber.org/zap to v1.24.0 ([#268](https://github.com/james-milligan/open-feature-operator/issues/268)) ([b7bdde8](https://github.com/james-milligan/open-feature-operator/commit/b7bdde8944446621751e6ef70e6b0f0646adee21))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.13.1 ([edeffcd](https://github.com/james-milligan/open-feature-operator/commit/edeffcd3ef6fe9a8d52d0d5c414512ef8cd80629))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.5 ([#279](https://github.com/james-milligan/open-feature-operator/issues/279)) ([8a80bff](https://github.com/james-milligan/open-feature-operator/commit/8a80bff886af404e897e6a247cea2f4c88d88499))
* for helm issues ([#206](https://github.com/james-milligan/open-feature-operator/issues/206)) ([39febd7](https://github.com/james-milligan/open-feature-operator/commit/39febd76d1b996afdbc24399bcd08b502621c6cc))
* gave configmaps volume mounts a subpath to allow for multiple mounts ([#321](https://github.com/james-milligan/open-feature-operator/issues/321)) ([2ec454c](https://github.com/james-milligan/open-feature-operator/commit/2ec454c036149ebeaf34f81cbf4ad7895f0bb995))
* include assets in release ([#109](https://github.com/james-milligan/open-feature-operator/issues/109)) ([b835abb](https://github.com/james-milligan/open-feature-operator/commit/b835abb48ae8ca3c9c63abd51ae5614a4068c003))
* include release tag in helm charts publishing ([2746716](https://github.com/james-milligan/open-feature-operator/commit/27467164dcd05b0220e0857bf79e42d62e7a40a9))
* mount dirs not files ([#326](https://github.com/james-milligan/open-feature-operator/issues/326)) ([089ab3c](https://github.com/james-milligan/open-feature-operator/commit/089ab3c48c0937e64060057e43ff07cf8fd47f67))
* nil pointer dereference ([#216](https://github.com/james-milligan/open-feature-operator/issues/216)) ([d975066](https://github.com/james-milligan/open-feature-operator/commit/d975066f96a5f9caf8af8d513076480a33943257))
* package signing fixes ([36597f4](https://github.com/james-milligan/open-feature-operator/commit/36597f484c85effd6a993f44b97fcd541d34c515))
* Package signing should happen in the oci workflow. ([a04a110](https://github.com/james-milligan/open-feature-operator/commit/a04a110e29b1725a66d0f4b529741947ebb7c798))
* remove redundant name ([#189](https://github.com/james-milligan/open-feature-operator/issues/189)) ([664bd73](https://github.com/james-milligan/open-feature-operator/commit/664bd7314e376b23a01247b5c027c04a9ac26329))
* removed duplicate config map generation, resolve permissions issue ([#305](https://github.com/james-milligan/open-feature-operator/issues/305)) ([eec16af](https://github.com/james-milligan/open-feature-operator/commit/eec16af28eb963a3d0f276d382e808079e663a50))
* resolve issue with templated DNS name in cert ([65068df](https://github.com/james-milligan/open-feature-operator/commit/65068df3019312a965271e50c52bbb90b68665c0))
* Security issues ([#348](https://github.com/james-milligan/open-feature-operator/issues/348)) ([5bd0b19](https://github.com/james-milligan/open-feature-operator/commit/5bd0b192a5db7f1557e1161e4bb425bbf0e31e2a))
* set defaultTag to INPUT_FLAGD_VERSION ([#332](https://github.com/james-milligan/open-feature-operator/issues/332)) ([23547a1](https://github.com/james-milligan/open-feature-operator/commit/23547a1e155e0cde2f085882bfd43128681466cd))
* set sbom dir ([#178](https://github.com/james-milligan/open-feature-operator/issues/178)) ([143adf9](https://github.com/james-milligan/open-feature-operator/commit/143adf910fe15a8b8af31dff48743352ab203d83))
* set sbom dir ([#180](https://github.com/james-milligan/open-feature-operator/issues/180)) ([616272d](https://github.com/james-milligan/open-feature-operator/commit/616272d6d693115a22839cf52eb8fd448609ad6c))
* uniqueness of featureflagconfiguration file path ([#323](https://github.com/james-milligan/open-feature-operator/issues/323)) ([2b10945](https://github.com/james-milligan/open-feature-operator/commit/2b109452893abd053640ffbb9c79b834b78feb7b))
* update flagd version ([#402](https://github.com/james-milligan/open-feature-operator/issues/402)) ([dc6aa3c](https://github.com/james-milligan/open-feature-operator/commit/dc6aa3c3dd9fec6c508b34608384247b63b42eeb))
* update x/net for CVE-2022-41721 ([#301](https://github.com/james-milligan/open-feature-operator/issues/301)) ([bbe9837](https://github.com/james-milligan/open-feature-operator/commit/bbe983786ff74b59046b95082d79f71089fe2b67))
* Upload sbom ([#175](https://github.com/james-milligan/open-feature-operator/issues/175)) ([813c646](https://github.com/james-milligan/open-feature-operator/commit/813c6469ecc18101f60c593282ed32d7579f5880))
* Upload sbom by name ([#176](https://github.com/james-milligan/open-feature-operator/issues/176)) ([7d0fcd0](https://github.com/james-milligan/open-feature-operator/commit/7d0fcd0ba7eeee1b2424189c7e5f5f92bc1fffac))
* Version fix ([#284](https://github.com/james-milligan/open-feature-operator/issues/284)) ([a9c6f15](https://github.com/james-milligan/open-feature-operator/commit/a9c6f154589f1e00e60883c229b3ee29d7d2e9aa))


### 🧹 Chore

* adapt Makefile for local dev deployment + introduce report logs for e2e-tests ([#359](https://github.com/james-milligan/open-feature-operator/issues/359)) ([17ac016](https://github.com/james-milligan/open-feature-operator/commit/17ac016a3f326c726662c55dd19daed8777c76a1))
* add additional Kubernetes files ([#258](https://github.com/james-milligan/open-feature-operator/issues/258)) ([b04928b](https://github.com/james-milligan/open-feature-operator/commit/b04928b5d96196f6590f46e37a413d74d3186780))
* add additional sections to the release notes ([4bec5af](https://github.com/james-milligan/open-feature-operator/commit/4bec5af5fc5fc589d920f0c17a1213a036b558a0))
* add all resources individually to tarball ([#345](https://github.com/james-milligan/open-feature-operator/issues/345)) ([a0993dc](https://github.com/james-milligan/open-feature-operator/commit/a0993dc5a2e1f3550aab2656b9ddeb08674dc896))
* add artifact hub metadata ([#372](https://github.com/james-milligan/open-feature-operator/issues/372)) ([c6f539f](https://github.com/james-milligan/open-feature-operator/commit/c6f539f5bdd9dc18ac197eb3303d91131e863011))
* Add beeme1mr as a code owner ([99472c7](https://github.com/james-milligan/open-feature-operator/commit/99472c784350df90940ed09243bf51ee4fa60370))
* add latest tag ([#125](https://github.com/james-milligan/open-feature-operator/issues/125)) ([33474fe](https://github.com/james-milligan/open-feature-operator/commit/33474fe3f643d4924207bd1ce28143b5dd996a1b))
* added cert-manager installation command to helm charts installation section ([2d4837d](https://github.com/james-milligan/open-feature-operator/commit/2d4837d7e237f52962e5f025ca394ce1ea48658d))
* added crd configuration reconciliation e2e test ([f4a1c25](https://github.com/james-milligan/open-feature-operator/commit/f4a1c256312026cb0fc7d5986d2e0ccddbf83ee9))
* automate release of charts ([552d345](https://github.com/james-milligan/open-feature-operator/commit/552d34550c83ebaa06174c1c0ee34c0cdc1dd747))
* bump ENVTEST_K8S_VERSION to 1.25 ([#214](https://github.com/james-milligan/open-feature-operator/issues/214)) ([a3d0609](https://github.com/james-milligan/open-feature-operator/commit/a3d06090544310519e8574ca81873840082b6acf))
* bump go version in release-assets workflow ([#230](https://github.com/james-milligan/open-feature-operator/issues/230)) ([8d4d587](https://github.com/james-milligan/open-feature-operator/commit/8d4d58768fe194c5c99010ca9ee697ef9ebc84ae))
* **deps:** add renovate.json ([c9450f1](https://github.com/james-milligan/open-feature-operator/commit/c9450f1717af67ad5a62b12a3c1f7237f70bb3b4))
* **deps:** update actions/checkout action to v3 ([b56db1b](https://github.com/james-milligan/open-feature-operator/commit/b56db1b7c03555bff76e1f0f96bbd6e427179e4a))
* **deps:** update amannn/action-semantic-pull-request action to v5 ([295dc91](https://github.com/james-milligan/open-feature-operator/commit/295dc9157ff8345cbdd8f0d0d1e078c26af232f4))
* **deps:** update dependency open-feature/flagd to v0.3.1 ([#296](https://github.com/james-milligan/open-feature-operator/issues/296)) ([1eff914](https://github.com/james-milligan/open-feature-operator/commit/1eff914228d5ca04b9781db1e0ead1dd7aef0462))
* **deps:** update dependency open-feature/flagd to v0.3.2 ([#314](https://github.com/james-milligan/open-feature-operator/issues/314)) ([dbfea97](https://github.com/james-milligan/open-feature-operator/commit/dbfea97ac471faee34e168552dae2be6d2416b9c))
* **deps:** update dependency open-feature/flagd to v0.3.4 ([#327](https://github.com/james-milligan/open-feature-operator/issues/327)) ([eb4d6ee](https://github.com/james-milligan/open-feature-operator/commit/eb4d6ee4e7d8a29f8eb25fc649a93373eab4aca7))
* **deps:** update dependency open-feature/flagd to v0.4.0 ([#342](https://github.com/james-milligan/open-feature-operator/issues/342)) ([0640f46](https://github.com/james-milligan/open-feature-operator/commit/0640f469daa3c0adce920bb73e901fe83bc275e7))
* **deps:** update dependency open-feature/flagd to v0.4.1 ([#373](https://github.com/james-milligan/open-feature-operator/issues/373)) ([756cf7a](https://github.com/james-milligan/open-feature-operator/commit/756cf7a96c05fdfa8ffa2bf933225b84af400e37))
* **deps:** update dependency open-feature/flagd to v0.4.4 ([#400](https://github.com/james-milligan/open-feature-operator/issues/400)) ([3e0a666](https://github.com/james-milligan/open-feature-operator/commit/3e0a666f2824071c49250a4467d62b96a5af5ee7))
* **deps:** update docker/build-push-action action to v3 ([214bf57](https://github.com/james-milligan/open-feature-operator/commit/214bf57f89de6e377d17d5f8f6c801fdfb2d7061))
* **deps:** update docker/build-push-action action to v4 ([#335](https://github.com/james-milligan/open-feature-operator/issues/335)) ([a4cfad8](https://github.com/james-milligan/open-feature-operator/commit/a4cfad87193c624fbc9cab5235fbea0adefdbb90))
* **deps:** update docker/login-action digest to 219c305 ([#365](https://github.com/james-milligan/open-feature-operator/issues/365)) ([ee84954](https://github.com/james-milligan/open-feature-operator/commit/ee849546322516019ea19a205c22c4ee38ac36ed))
* **deps:** update docker/login-action digest to 3da7dc6 ([#266](https://github.com/james-milligan/open-feature-operator/issues/266)) ([7e8ed83](https://github.com/james-milligan/open-feature-operator/commit/7e8ed830825b78cc9053ade134da979e44eb243d))
* **deps:** update docker/login-action digest to f4ef78c ([71b37aa](https://github.com/james-milligan/open-feature-operator/commit/71b37aa5e5928234bf403afa123b3665f64c294c))
* **deps:** update docker/metadata-action digest to 5739616 ([7b82bd0](https://github.com/james-milligan/open-feature-operator/commit/7b82bd0a2a2c1119029c97d930f3b6606de2059b))
* **deps:** update docker/metadata-action digest to 766400c ([#267](https://github.com/james-milligan/open-feature-operator/issues/267)) ([38a1464](https://github.com/james-milligan/open-feature-operator/commit/38a14644e687b928e51d1350f6d57ef9d493330c))
* **deps:** update docker/metadata-action digest to 9ec57ed ([#366](https://github.com/james-milligan/open-feature-operator/issues/366)) ([884d444](https://github.com/james-milligan/open-feature-operator/commit/884d44422ad7bfa28a8fb88156cd66e252e2eba5))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.13.0 ([86953f6](https://github.com/james-milligan/open-feature-operator/commit/86953f63c691cd4cce56854de46022a7e898f82c))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.13.1 ([2462b1c](https://github.com/james-milligan/open-feature-operator/commit/2462b1c2b298927bea0d43eb03ee885e6d43c4c9))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.13.1 ([31570b2](https://github.com/james-milligan/open-feature-operator/commit/31570b2041ab143d6e1e6cdce90f03c449b03e17))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.14.0 ([#376](https://github.com/james-milligan/open-feature-operator/issues/376)) ([708e4bc](https://github.com/james-milligan/open-feature-operator/commit/708e4bc44d8493d4f0aaa7f7036c2b7ecd2efd32))
* **deps:** update ghcr.io/open-feature/flagd docker tag to v0.3.2 ([#315](https://github.com/james-milligan/open-feature-operator/issues/315)) ([c178c49](https://github.com/james-milligan/open-feature-operator/commit/c178c49f585e6e8ff28d8dff49c6c95930dd4cef))
* **deps:** update ghcr.io/open-feature/flagd docker tag to v0.3.4 ([#328](https://github.com/james-milligan/open-feature-operator/issues/328)) ([f96246f](https://github.com/james-milligan/open-feature-operator/commit/f96246f748b00094bfb82b6b0dcf035a17128dc9))
* **deps:** update ghcr.io/open-feature/flagd docker tag to v0.4.4 ([#381](https://github.com/james-milligan/open-feature-operator/issues/381)) ([a253761](https://github.com/james-milligan/open-feature-operator/commit/a253761af8565fdcf6e6f9ca92c740f25b4b0620))
* **deps:** update ghcr.io/open-feature/open-feature-operator docker tag to v0.2.16 ([7644166](https://github.com/james-milligan/open-feature-operator/commit/764416681b6516ab07bda146afa1d3e4ac142dba))
* **deps:** update golang docker tag to v1.20.2 ([#374](https://github.com/james-milligan/open-feature-operator/issues/374)) ([e2de529](https://github.com/james-milligan/open-feature-operator/commit/e2de52997b44835a4a8515e9fd37c976d3539272))
* **deps:** update helm/kind-action action to v1.5.0 ([#277](https://github.com/james-milligan/open-feature-operator/issues/277)) ([15ad251](https://github.com/james-milligan/open-feature-operator/commit/15ad251614020a4581de015ffc5c721101326ead))
* **docs:** add concepts doc, add filepath docs ([#340](https://github.com/james-milligan/open-feature-operator/issues/340)) ([5f83920](https://github.com/james-milligan/open-feature-operator/commit/5f839209008ced69198cdf909659c0001e1b3910))
* document crd upgrades ([#361](https://github.com/james-milligan/open-feature-operator/issues/361)) ([e708990](https://github.com/james-milligan/open-feature-operator/commit/e70899081b8c0864c8ef17f75275d67b26b63c20))
* document the end to end installation process with more explanation ([4aec1fc](https://github.com/james-milligan/open-feature-operator/commit/4aec1fca816c1e4130aa70cab49f8f4f406cad26))
* e2e test for openfeature.dev/enabled annotation set to false ([#375](https://github.com/james-milligan/open-feature-operator/issues/375)) ([b03fb14](https://github.com/james-milligan/open-feature-operator/commit/b03fb145e317f987727d76b98041fa783e5c2202))
* fix image name in helm chart ([f2782c6](https://github.com/james-milligan/open-feature-operator/commit/f2782c6d381fee59b37f3be3f5b1399a1ee7eb29))
* fix make helm-package command in release charts workflow ([3fc9735](https://github.com/james-milligan/open-feature-operator/commit/3fc9735148619771693bfbb2a92f62e12bfd9480))
* fix permissions for image push ([d4d63c2](https://github.com/james-milligan/open-feature-operator/commit/d4d63c20eca5dd0bb040104807e0bd5eee1ceb81))
* gitignore generate charts ([#356](https://github.com/james-milligan/open-feature-operator/issues/356)) ([250b5ab](https://github.com/james-milligan/open-feature-operator/commit/250b5abd6d6dedebf23bdf931777e53ec116565f))
* go version bump ([5219a1e](https://github.com/james-milligan/open-feature-operator/commit/5219a1e6a4d449bcdbae626f2a8cba5dab622426))
* how to implement a new crd version ([#227](https://github.com/james-milligan/open-feature-operator/issues/227)) ([1b01cec](https://github.com/james-milligan/open-feature-operator/commit/1b01cec8e81b1a6dd9da0f5a94871eafa934d2a9))
* improve formatting and content ([#384](https://github.com/james-milligan/open-feature-operator/issues/384)) ([c5a6a32](https://github.com/james-milligan/open-feature-operator/commit/c5a6a32f0ccccc6449fc581de08c283434c1adb6))
* integration tests for pod mutation webhook & featureflagconfiguration validation webhook ([2295659](https://github.com/james-milligan/open-feature-operator/commit/2295659b47dfff46a0690b051db6a78c368b576c))
* **main:** release 0.2.1 ([#112](https://github.com/james-milligan/open-feature-operator/issues/112)) ([0fca997](https://github.com/james-milligan/open-feature-operator/commit/0fca99765a4954075123d441b33a8f746e3ac2e6))
* **main:** release 0.2.10 ([29bb985](https://github.com/james-milligan/open-feature-operator/commit/29bb985f3f6dfdab6d564bcbf729cb93c26db74b))
* **main:** release 0.2.11 ([#177](https://github.com/james-milligan/open-feature-operator/issues/177)) ([8cd768e](https://github.com/james-milligan/open-feature-operator/commit/8cd768e342266d766c7b85badcb015c76a4ed77a))
* **main:** release 0.2.12 ([#179](https://github.com/james-milligan/open-feature-operator/issues/179)) ([fe3ca2d](https://github.com/james-milligan/open-feature-operator/commit/fe3ca2d3e28a7044c8ff5ff1481ea20a3cd914e1))
* **main:** release 0.2.13 ([#181](https://github.com/james-milligan/open-feature-operator/issues/181)) ([e685eae](https://github.com/james-milligan/open-feature-operator/commit/e685eae3c84df53d890c9ab009c4f10e3ca14a36))
* **main:** release 0.2.14 ([#183](https://github.com/james-milligan/open-feature-operator/issues/183)) ([0145eca](https://github.com/james-milligan/open-feature-operator/commit/0145eca05faa525800a96db01c3e3833443eccf3))
* **main:** release 0.2.15 ([#188](https://github.com/james-milligan/open-feature-operator/issues/188)) ([0068511](https://github.com/james-milligan/open-feature-operator/commit/0068511e46b093dc7064c8026302feab7c2a3fbf))
* **main:** release 0.2.16 ([59af227](https://github.com/james-milligan/open-feature-operator/commit/59af227f36586dff5d4d3943cefa7038ab7540a5))
* **main:** release 0.2.17 ([9546195](https://github.com/james-milligan/open-feature-operator/commit/9546195d0d30f374780e72e59a39c01ac045d5c3))
* **main:** release 0.2.17 ([#208](https://github.com/james-milligan/open-feature-operator/issues/208)) ([f375eb2](https://github.com/james-milligan/open-feature-operator/commit/f375eb2232f588931c7391367d8cfb58f026eb89))
* **main:** release 0.2.18 ([#218](https://github.com/james-milligan/open-feature-operator/issues/218)) ([d4be9ae](https://github.com/james-milligan/open-feature-operator/commit/d4be9ae7280a7a5c0448a3efa23f59bf5d05096c))
* **main:** release 0.2.19 ([12b3934](https://github.com/james-milligan/open-feature-operator/commit/12b3934ae20874ffe822036d4e0242bf95ce09c2))
* **main:** release 0.2.2 ([#115](https://github.com/james-milligan/open-feature-operator/issues/115)) ([b1ed770](https://github.com/james-milligan/open-feature-operator/commit/b1ed770c572d072b046d32f5d1de226bcf13ac33))
* **main:** release 0.2.20 ([d22af86](https://github.com/james-milligan/open-feature-operator/commit/d22af86e258a2ffa85c986a4898da9f6e27fb2e5))
* **main:** release 0.2.21 ([df8059f](https://github.com/james-milligan/open-feature-operator/commit/df8059f183fda3aa063df8c6d054163c0dd3cfb1))
* **main:** release 0.2.22 ([#283](https://github.com/james-milligan/open-feature-operator/issues/283)) ([40f5c0e](https://github.com/james-milligan/open-feature-operator/commit/40f5c0e4135e19ba152285f43bd2627ea7c0be73))
* **main:** release 0.2.23 ([#290](https://github.com/james-milligan/open-feature-operator/issues/290)) ([53ef2b9](https://github.com/james-milligan/open-feature-operator/commit/53ef2b95a5b829ce30d10b9ae535842511169a5f))
* **main:** release 0.2.24 ([#300](https://github.com/james-milligan/open-feature-operator/issues/300)) ([6a78925](https://github.com/james-milligan/open-feature-operator/commit/6a78925e2fe852a02845acc32309784bb0dbf816))
* **main:** release 0.2.25 ([#306](https://github.com/james-milligan/open-feature-operator/issues/306)) ([27f7d4d](https://github.com/james-milligan/open-feature-operator/commit/27f7d4df57ad23da45eaeafc94a37b7f392f3c0e))
* **main:** release 0.2.26 ([#316](https://github.com/james-milligan/open-feature-operator/issues/316)) ([d7e9559](https://github.com/james-milligan/open-feature-operator/commit/d7e9559a9f60b2eab5a6e86fea734c04ad05980a))
* **main:** release 0.2.27 ([#322](https://github.com/james-milligan/open-feature-operator/issues/322)) ([256894f](https://github.com/james-milligan/open-feature-operator/commit/256894fe52797cef22b7186d354ec6056224883a))
* **main:** release 0.2.28 ([#330](https://github.com/james-milligan/open-feature-operator/issues/330)) ([312e91e](https://github.com/james-milligan/open-feature-operator/commit/312e91e6f9c1c44b9d642c74e031742f78c0f7f9))
* **main:** release 0.2.29 ([#339](https://github.com/james-milligan/open-feature-operator/issues/339)) ([8b779b1](https://github.com/james-milligan/open-feature-operator/commit/8b779b1855de171b0f81197d8850dbd2dff9322b))
* **main:** release 0.2.3 ([#122](https://github.com/james-milligan/open-feature-operator/issues/122)) ([da1635a](https://github.com/james-milligan/open-feature-operator/commit/da1635a0e69ee4d257ae7066ca8f80f0557d8080))
* **main:** release 0.2.30 ([#363](https://github.com/james-milligan/open-feature-operator/issues/363)) ([945aa6b](https://github.com/james-milligan/open-feature-operator/commit/945aa6b81d27803a37771b90a08f5a7a79bb8085))
* **main:** release 0.2.31 ([#403](https://github.com/james-milligan/open-feature-operator/issues/403)) ([42f9b00](https://github.com/james-milligan/open-feature-operator/commit/42f9b007d391137b01bf9bc36bc40666d16c18b3))
* **main:** release 0.2.4 ([#124](https://github.com/james-milligan/open-feature-operator/issues/124)) ([1de6abc](https://github.com/james-milligan/open-feature-operator/commit/1de6abc31990633ab386ed5fe604a0408fd7c493))
* **main:** release 0.2.5 ([#127](https://github.com/james-milligan/open-feature-operator/issues/127)) ([faaf923](https://github.com/james-milligan/open-feature-operator/commit/faaf9235623eb9afbd50ee24f1004419bf38ed8e))
* **main:** release 0.2.6 ([5a705a1](https://github.com/james-milligan/open-feature-operator/commit/5a705a1063a472b1e3989d97a62c013be0770db0))
* **main:** release 0.2.6 ([#136](https://github.com/james-milligan/open-feature-operator/issues/136)) ([80bcfba](https://github.com/james-milligan/open-feature-operator/commit/80bcfbae8475badb1574f17a4a3bfe1b3ef11088))
* **main:** release 0.2.7 ([73d9fac](https://github.com/james-milligan/open-feature-operator/commit/73d9fac4f57bc0114deba58426016b28fcd7045e))
* **main:** release 0.2.7 ([baf50be](https://github.com/james-milligan/open-feature-operator/commit/baf50be205c88a862b4745994cf2775389278ab1))
* **main:** release 0.2.8 ([5045838](https://github.com/james-milligan/open-feature-operator/commit/50458381ab8fe7c45c42698f88fa2125fec0c93d))
* **main:** release 0.2.9 ([237e951](https://github.com/james-milligan/open-feature-operator/commit/237e951ac54e7fbc5c7b04f310658e62802aba47))
* remove namespace argument from e2e example in readme ([#217](https://github.com/james-milligan/open-feature-operator/issues/217)) ([cb3b06d](https://github.com/james-milligan/open-feature-operator/commit/cb3b06dd41c71de332faa93008a4425069dabdd2))
* remove pre alpha badges ([#317](https://github.com/james-milligan/open-feature-operator/issues/317)) ([5fcf637](https://github.com/james-milligan/open-feature-operator/commit/5fcf637dc6a1a18bb1234ec8a621195a8c71551e))
* remove published charts ([9dafe03](https://github.com/james-milligan/open-feature-operator/commit/9dafe03a115fec6c1d2b998d0b6d98c00c5586cb))
* remove unneeded conversion webhooks + introduce unit tests for conversion functions ([#385](https://github.com/james-milligan/open-feature-operator/issues/385)) ([dd34801](https://github.com/james-milligan/open-feature-operator/commit/dd34801fd71ac4f1e6c0b0f39f78ddf738f5601d))
* renovate - force semantic commits ([152bf59](https://github.com/james-milligan/open-feature-operator/commit/152bf59d4e5edf6ad6fa304f12093d12e0293491))
* renovate regex manager to monitor FLAGD_VERSION in Makefile ([e7f5e4b](https://github.com/james-milligan/open-feature-operator/commit/e7f5e4b5eaf0107b01a04f71bb57791e1ba7a3ee))
* revert kustomization image name change ([#232](https://github.com/james-milligan/open-feature-operator/issues/232)) ([0956a00](https://github.com/james-milligan/open-feature-operator/commit/0956a0054003ba57c8bd890a3bdf4030a42a6607))
* run e2e test as a part of the github workflow ([99bb25c](https://github.com/james-milligan/open-feature-operator/commit/99bb25cc134f3d967eea3fc5ea7ebc7c0467d99f))
* test filepath sync config change is reconciliated ([#338](https://github.com/james-milligan/open-feature-operator/issues/338)) ([101729a](https://github.com/james-milligan/open-feature-operator/commit/101729ad587d63541227c7eb9088ff67bf412640))
* test multiple sync providers ([#325](https://github.com/james-milligan/open-feature-operator/issues/325)) ([f3fe057](https://github.com/james-milligan/open-feature-operator/commit/f3fe057e180528944110a6841239d3f9d471086b))
* threat modeling ([#309](https://github.com/james-milligan/open-feature-operator/issues/309)) ([dc7d935](https://github.com/james-milligan/open-feature-operator/commit/dc7d9351b7f75bcc8e8eda3c3ea1d72e145955be))
* update build badge to new format ([13e85de](https://github.com/james-milligan/open-feature-operator/commit/13e85de7764b602f8b08687272f6b6483b6d13f0))
* update charts to v0.2.18 ([d5fc087](https://github.com/james-milligan/open-feature-operator/commit/d5fc08762dab469ecd4f62219d2e0333c507dd46))
* update charts to v0.2.19 ([fbd82eb](https://github.com/james-milligan/open-feature-operator/commit/fbd82eb56ba322b161476cebbc881f2535b1527f))
* update CODEOWNERS ([edb1592](https://github.com/james-milligan/open-feature-operator/commit/edb15923ff8b029f05553296df4bb135b8f697d2))
* Updated workflow permissions for CLOMonitor status ([#138](https://github.com/james-milligan/open-feature-operator/issues/138)) ([8733065](https://github.com/james-milligan/open-feature-operator/commit/8733065ed3db59d07d4f853f4b529949f5ce69a1))
* upgrade ENVTEST_K8S_VERSION to 1.26.1 ([#360](https://github.com/james-milligan/open-feature-operator/issues/360)) ([5e01c24](https://github.com/james-milligan/open-feature-operator/commit/5e01c249a0da46eea0f964f2ded69a4e20c10335))
* use renovate to bump flagd version ([#395](https://github.com/james-milligan/open-feature-operator/issues/395)) ([fd5b072](https://github.com/james-milligan/open-feature-operator/commit/fd5b072214f1c3c74dfc4bc53ca1ff6c14d72ffa))

## [0.2.31](https://github.com/open-feature/open-feature-operator/compare/v0.2.30...v0.2.31) (2023-03-16)


### 📚 Documentation

* fix rendering issue with operator resource config table ([#401](https://github.com/open-feature/open-feature-operator/issues/401)) ([71ea8a6](https://github.com/open-feature/open-feature-operator/commit/71ea8a68bbb97052822552ffce3c498c3da0e52d))


### 🐛 Bug Fixes

* update flagd version ([#402](https://github.com/open-feature/open-feature-operator/issues/402)) ([dc6aa3c](https://github.com/open-feature/open-feature-operator/commit/dc6aa3c3dd9fec6c508b34608384247b63b42eeb))

## [0.2.30](https://github.com/open-feature/open-feature-operator/compare/v0.2.29...v0.2.30) (2023-03-16)


### 📚 Documentation

* add AND operator to sequential commands ([#368](https://github.com/open-feature/open-feature-operator/issues/368)) ([6f73a62](https://github.com/open-feature/open-feature-operator/commit/6f73a6214d87771f9555469fe4d60dbb6d301198))


### ✨ New Features

* enable flagd probes ([#390](https://github.com/open-feature/open-feature-operator/issues/390)) ([41efb15](https://github.com/open-feature/open-feature-operator/commit/41efb155994b3cfb768cc39e59bfc09781c57f2e))
* improve deployment pattern ([#344](https://github.com/open-feature/open-feature-operator/issues/344)) ([572ba96](https://github.com/open-feature/open-feature-operator/commit/572ba961912ada2c07eb6143925d16ab6a6a85a3))


### 🐛 Bug Fixes

* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.5 ([#279](https://github.com/open-feature/open-feature-operator/issues/279)) ([8a80bff](https://github.com/open-feature/open-feature-operator/commit/8a80bff886af404e897e6a247cea2f4c88d88499))


### 🧹 Chore

* add additional sections to the release notes ([4bec5af](https://github.com/open-feature/open-feature-operator/commit/4bec5af5fc5fc589d920f0c17a1213a036b558a0))
* add artifact hub metadata ([#372](https://github.com/open-feature/open-feature-operator/issues/372)) ([c6f539f](https://github.com/open-feature/open-feature-operator/commit/c6f539f5bdd9dc18ac197eb3303d91131e863011))
* **deps:** update dependency open-feature/flagd to v0.4.0 ([#342](https://github.com/open-feature/open-feature-operator/issues/342)) ([0640f46](https://github.com/open-feature/open-feature-operator/commit/0640f469daa3c0adce920bb73e901fe83bc275e7))
* **deps:** update dependency open-feature/flagd to v0.4.1 ([#373](https://github.com/open-feature/open-feature-operator/issues/373)) ([756cf7a](https://github.com/open-feature/open-feature-operator/commit/756cf7a96c05fdfa8ffa2bf933225b84af400e37))
* **deps:** update dependency open-feature/flagd to v0.4.4 ([#400](https://github.com/open-feature/open-feature-operator/issues/400)) ([3e0a666](https://github.com/open-feature/open-feature-operator/commit/3e0a666f2824071c49250a4467d62b96a5af5ee7))
* **deps:** update docker/login-action digest to 219c305 ([#365](https://github.com/open-feature/open-feature-operator/issues/365)) ([ee84954](https://github.com/open-feature/open-feature-operator/commit/ee849546322516019ea19a205c22c4ee38ac36ed))
* **deps:** update docker/metadata-action digest to 766400c ([#267](https://github.com/open-feature/open-feature-operator/issues/267)) ([38a1464](https://github.com/open-feature/open-feature-operator/commit/38a14644e687b928e51d1350f6d57ef9d493330c))
* **deps:** update docker/metadata-action digest to 9ec57ed ([#366](https://github.com/open-feature/open-feature-operator/issues/366)) ([884d444](https://github.com/open-feature/open-feature-operator/commit/884d44422ad7bfa28a8fb88156cd66e252e2eba5))
* **deps:** update gcr.io/kubebuilder/kube-rbac-proxy docker tag to v0.14.0 ([#376](https://github.com/open-feature/open-feature-operator/issues/376)) ([708e4bc](https://github.com/open-feature/open-feature-operator/commit/708e4bc44d8493d4f0aaa7f7036c2b7ecd2efd32))
* **deps:** update ghcr.io/open-feature/flagd docker tag to v0.4.4 ([#381](https://github.com/open-feature/open-feature-operator/issues/381)) ([a253761](https://github.com/open-feature/open-feature-operator/commit/a253761af8565fdcf6e6f9ca92c740f25b4b0620))
* **deps:** update golang docker tag to v1.20.2 ([#374](https://github.com/open-feature/open-feature-operator/issues/374)) ([e2de529](https://github.com/open-feature/open-feature-operator/commit/e2de52997b44835a4a8515e9fd37c976d3539272))
* e2e test for openfeature.dev/enabled annotation set to false ([#375](https://github.com/open-feature/open-feature-operator/issues/375)) ([b03fb14](https://github.com/open-feature/open-feature-operator/commit/b03fb145e317f987727d76b98041fa783e5c2202))
* improve formatting and content ([#384](https://github.com/open-feature/open-feature-operator/issues/384)) ([c5a6a32](https://github.com/open-feature/open-feature-operator/commit/c5a6a32f0ccccc6449fc581de08c283434c1adb6))
* remove unneeded conversion webhooks + introduce unit tests for conversion functions ([#385](https://github.com/open-feature/open-feature-operator/issues/385)) ([dd34801](https://github.com/open-feature/open-feature-operator/commit/dd34801fd71ac4f1e6c0b0f39f78ddf738f5601d))

## [0.2.29](https://github.com/open-feature/open-feature-operator/compare/v0.2.28...v0.2.29) (2023-02-23)


### Features

* add log format configuration options through helm chart ([#346](https://github.com/open-feature/open-feature-operator/issues/346)) ([bcef736](https://github.com/open-feature/open-feature-operator/commit/bcef7368fc4905b351f81f5dfa10eb4c26bf8764))
* Introduced context to the readyz endpoint, added wait to test suite ([#336](https://github.com/open-feature/open-feature-operator/issues/336)) ([ed81c02](https://github.com/open-feature/open-feature-operator/commit/ed81c0284f8d759eb228d3af7030efb0b94ee280))


### Bug Fixes

* Security issues ([#348](https://github.com/open-feature/open-feature-operator/issues/348)) ([5bd0b19](https://github.com/open-feature/open-feature-operator/commit/5bd0b192a5db7f1557e1161e4bb425bbf0e31e2a))
* set defaultTag to INPUT_FLAGD_VERSION ([#332](https://github.com/open-feature/open-feature-operator/issues/332)) ([23547a1](https://github.com/open-feature/open-feature-operator/commit/23547a1e155e0cde2f085882bfd43128681466cd))

## [0.2.28](https://github.com/open-feature/open-feature-operator/compare/v0.2.27...v0.2.28) (2023-01-28)


### Bug Fixes

* mount dirs not files ([#326](https://github.com/open-feature/open-feature-operator/issues/326)) ([089ab3c](https://github.com/open-feature/open-feature-operator/commit/089ab3c48c0937e64060057e43ff07cf8fd47f67))

## [0.2.27](https://github.com/open-feature/open-feature-operator/compare/v0.2.26...v0.2.27) (2023-01-27)


### Features

* default sync provider configuration ([#320](https://github.com/open-feature/open-feature-operator/issues/320)) ([7cba7e1](https://github.com/open-feature/open-feature-operator/commit/7cba7e14c223a083f02ff8313b899583253120f3))


### Bug Fixes

* gave configmaps volume mounts a subpath to allow for multiple mounts ([#321](https://github.com/open-feature/open-feature-operator/issues/321)) ([2ec454c](https://github.com/open-feature/open-feature-operator/commit/2ec454c036149ebeaf34f81cbf4ad7895f0bb995))
* uniqueness of featureflagconfiguration file path ([#323](https://github.com/open-feature/open-feature-operator/issues/323)) ([2b10945](https://github.com/open-feature/open-feature-operator/commit/2b109452893abd053640ffbb9c79b834b78feb7b))

## [0.2.26](https://github.com/open-feature/open-feature-operator/compare/v0.2.25...v0.2.26) (2023-01-26)


### Bug Fixes

* **deps:** update module github.com/open-feature/schemas to v0.2.8 ([#269](https://github.com/open-feature/open-feature-operator/issues/269)) ([ed48060](https://github.com/open-feature/open-feature-operator/commit/ed48060b1f9e591ddadca4f9478728a823e10685))

## [0.2.25](https://github.com/open-feature/open-feature-operator/compare/v0.2.24...v0.2.25) (2023-01-25)


### Features

* Helm configuration ([#304](https://github.com/open-feature/open-feature-operator/issues/304)) ([99edfeb](https://github.com/open-feature/open-feature-operator/commit/99edfeb8c32ada435f830c6799540ebdf3b5fcdd))


### Bug Fixes

* removed duplicate config map generation, resolve permissions issue ([#305](https://github.com/open-feature/open-feature-operator/issues/305)) ([eec16af](https://github.com/open-feature/open-feature-operator/commit/eec16af28eb963a3d0f276d382e808079e663a50))
* update x/net for CVE-2022-41721 ([#301](https://github.com/open-feature/open-feature-operator/issues/301)) ([bbe9837](https://github.com/open-feature/open-feature-operator/commit/bbe983786ff74b59046b95082d79f71089fe2b67))

## [0.2.24](https://github.com/open-feature/open-feature-operator/compare/v0.2.23...v0.2.24) (2023-01-16)


### Features

* backfill flagd-kubernetes-sync cluster role binding on startup ([#295](https://github.com/open-feature/open-feature-operator/pull/295))
* decouple feature flag spec from flagd config ([#276](https://github.com/open-feature/open-feature-operator/pull/276))


### Features

* upgrade flagd to v0.3.0 ([20571e1](https://github.com/open-feature/open-feature-operator/commit/20571e1018e102ffbcf01b2518fcbf8b66a287be))

## [0.2.22](https://github.com/open-feature/open-feature-operator/compare/v0.2.21...v0.2.22) (2022-12-16)


### Bug Fixes

* **deps:** update module go.uber.org/zap to v1.24.0 ([#268](https://github.com/open-feature/open-feature-operator/issues/268)) ([b7bdde8](https://github.com/open-feature/open-feature-operator/commit/b7bdde8944446621751e6ef70e6b0f0646adee21))
* Version fix ([#284](https://github.com/open-feature/open-feature-operator/issues/284)) ([a9c6f15](https://github.com/open-feature/open-feature-operator/commit/a9c6f154589f1e00e60883c229b3ee29d7d2e9aa))

## [0.2.21](https://github.com/open-feature/open-feature-operator/compare/v0.2.20...v0.2.21) (2022-12-16)


### Features

* add ff shortname, commit httpSyncConfiguration ([11e4652](https://github.com/open-feature/open-feature-operator/commit/11e46528fcd06cdc0c8e6f46944656224cd97441))
* introduce configurable resource limits for flagd sidecar ([e4affcf](https://github.com/open-feature/open-feature-operator/commit/e4affcfb0ccf13dc0406ef1c21c2b884a836f71f))


### Bug Fixes

* **deps:** update github.com/open-feature/schemas digest to 302d0fa ([#246](https://github.com/open-feature/open-feature-operator/issues/246)) ([7d22374](https://github.com/open-feature/open-feature-operator/commit/7d22374afb7a5e2e166550544d327ec7b5b3d1bf))
* **deps:** update kubernetes packages to v0.25.4 ([75bab2d](https://github.com/open-feature/open-feature-operator/commit/75bab2d441c945d51f17f0d32195a217072c3c15))
* include release tag in helm charts publishing ([2746716](https://github.com/open-feature/open-feature-operator/commit/27467164dcd05b0220e0857bf79e42d62e7a40a9))

## [0.2.20](https://github.com/open-feature/open-feature-operator/compare/v0.2.19...v0.2.20) (2022-11-18)


### Bug Fixes

* **deps:** update module sigs.k8s.io/controller-runtime to v0.13.1 ([edeffcd](https://github.com/open-feature/open-feature-operator/commit/edeffcd3ef6fe9a8d52d0d5c414512ef8cd80629))

## [0.2.19](https://github.com/open-feature/open-feature-operator/compare/v0.2.18...v0.2.19) (2022-11-15)


### Features

* introduced v1beta1 of featureflagconfiguration CRD with conversion webhook to v1alpha1 ([a45bdef](https://github.com/open-feature/open-feature-operator/commit/a45bdef5eec87738ce731af5825daffeb69eb6cb))
* structured the featureflagconfiguration CRD ([b056c7c](https://github.com/open-feature/open-feature-operator/commit/b056c7cdd76f4653c1a728342687beaa8279e314))

## [0.2.18](https://github.com/open-feature/open-feature-operator/compare/v0.2.17...v0.2.18) (2022-11-10)


### Bug Fixes

* nil pointer dereference ([#216](https://github.com/open-feature/open-feature-operator/issues/216)) ([d975066](https://github.com/open-feature/open-feature-operator/commit/d975066f96a5f9caf8af8d513076480a33943257))

## [0.2.17](https://github.com/open-feature/open-feature-operator/compare/v0.2.16...v0.2.17) (2022-11-07)


### Bug Fixes

* **deps:** update github.com/open-feature/schemas digest to d638ecf ([a984836](https://github.com/open-feature/open-feature-operator/commit/a98483696f467270783858046132f02b3d338ac2))
* for helm issues ([#206](https://github.com/open-feature/open-feature-operator/issues/206)) ([39febd7](https://github.com/open-feature/open-feature-operator/commit/39febd76d1b996afdbc24399bcd08b502621c6cc))

## [0.2.16](https://github.com/open-feature/open-feature-operator/compare/v0.2.15...v0.2.16) (2022-10-27)


### Bug Fixes

* resolve issue with templated DNS name in cert ([65068df](https://github.com/open-feature/open-feature-operator/commit/65068df3019312a965271e50c52bbb90b68665c0))

## [0.2.15](https://github.com/open-feature/open-feature-operator/compare/v0.2.14...v0.2.15) (2022-10-25)


### Bug Fixes

* artifact name and output file ([#187](https://github.com/open-feature/open-feature-operator/issues/187)) ([4dee157](https://github.com/open-feature/open-feature-operator/commit/4dee157d44c20fc925f9e33dbaae16c18f3d9b48))
* remove redundant name ([#189](https://github.com/open-feature/open-feature-operator/issues/189)) ([664bd73](https://github.com/open-feature/open-feature-operator/commit/664bd7314e376b23a01247b5c027c04a9ac26329))

## [0.2.14](https://github.com/open-feature/open-feature-operator/compare/v0.2.13...v0.2.14) (2022-10-25)


### Bug Fixes

* add sbom to ouput name ([#182](https://github.com/open-feature/open-feature-operator/issues/182)) ([5e939a8](https://github.com/open-feature/open-feature-operator/commit/5e939a8f67fbd095c18a6a2172bb856fe61dd173))

## [0.2.13](https://github.com/open-feature/open-feature-operator/compare/v0.2.12...v0.2.13) (2022-10-25)


### Bug Fixes

* set sbom dir ([#180](https://github.com/open-feature/open-feature-operator/issues/180)) ([616272d](https://github.com/open-feature/open-feature-operator/commit/616272d6d693115a22839cf52eb8fd448609ad6c))

## [0.2.12](https://github.com/open-feature/open-feature-operator/compare/v0.2.11...v0.2.12) (2022-10-25)


### Bug Fixes

* set sbom dir ([#178](https://github.com/open-feature/open-feature-operator/issues/178)) ([143adf9](https://github.com/open-feature/open-feature-operator/commit/143adf910fe15a8b8af31dff48743352ab203d83))

## [0.2.11](https://github.com/open-feature/open-feature-operator/compare/v0.2.10...v0.2.11) (2022-10-25)


### Bug Fixes

* Upload sbom ([#175](https://github.com/open-feature/open-feature-operator/issues/175)) ([813c646](https://github.com/open-feature/open-feature-operator/commit/813c6469ecc18101f60c593282ed32d7579f5880))
* Upload sbom by name ([#176](https://github.com/open-feature/open-feature-operator/issues/176)) ([7d0fcd0](https://github.com/open-feature/open-feature-operator/commit/7d0fcd0ba7eeee1b2424189c7e5f5f92bc1fffac))

## [0.2.10](https://github.com/open-feature/open-feature-operator/compare/v0.2.9...v0.2.10) (2022-10-25)


### Bug Fixes

* correcrt needs in asset release ([5ed4571](https://github.com/open-feature/open-feature-operator/commit/5ed45718ca189a15f7cdf4f8ddfc5864f189b1ce))

## [0.2.9](https://github.com/open-feature/open-feature-operator/compare/v0.2.8...v0.2.9) (2022-10-25)


### Bug Fixes

* Package signing should happen in the oci workflow. ([a04a110](https://github.com/open-feature/open-feature-operator/commit/a04a110e29b1725a66d0f4b529741947ebb7c798))

## [0.2.8](https://github.com/open-feature/open-feature-operator/compare/v0.2.7...v0.2.8) (2022-10-25)


### Bug Fixes

* package signing fixes ([36597f4](https://github.com/open-feature/open-feature-operator/commit/36597f484c85effd6a993f44b97fcd541d34c515))

## [0.2.7](https://github.com/open-feature/open-feature-operator/compare/v0.2.6...v0.2.7) (2022-10-25)


### Features

* adding artifacthub information ([#144](https://github.com/open-feature/open-feature-operator/issues/144)) ([65a5244](https://github.com/open-feature/open-feature-operator/commit/65a524445d1db8bb5608b88282a4d97a9bb6b74f))
* builds helm chart ([#137](https://github.com/open-feature/open-feature-operator/issues/137)) ([1525421](https://github.com/open-feature/open-feature-operator/commit/1525421229d43b17636dddb65d7b124e6477fe79))

## [0.2.7](https://github.com/open-feature/open-feature-operator/compare/v0.2.6...v0.2.7) (2022-10-24)


### Features

* adding artifacthub information ([#144](https://github.com/open-feature/open-feature-operator/issues/144)) ([65a5244](https://github.com/open-feature/open-feature-operator/commit/65a524445d1db8bb5608b88282a4d97a9bb6b74f))
* builds helm chart ([#137](https://github.com/open-feature/open-feature-operator/issues/137)) ([1525421](https://github.com/open-feature/open-feature-operator/commit/1525421229d43b17636dddb65d7b124e6477fe79))

## [0.2.6](https://github.com/open-feature/open-feature-operator/compare/v0.2.5...v0.2.6) (2022-10-24)


### Features

* adding artifacthub information ([#144](https://github.com/open-feature/open-feature-operator/issues/144)) ([65a5244](https://github.com/open-feature/open-feature-operator/commit/65a524445d1db8bb5608b88282a4d97a9bb6b74f))
* builds helm chart ([#137](https://github.com/open-feature/open-feature-operator/issues/137)) ([1525421](https://github.com/open-feature/open-feature-operator/commit/1525421229d43b17636dddb65d7b124e6477fe79))


### Bug Fixes

* CVE-2022-32149 ([015c19a](https://github.com/open-feature/open-feature-operator/commit/015c19ac4455673902c365111816b021f893c485))

## [0.2.6](https://github.com/open-feature/open-feature-operator/compare/v0.2.5...v0.2.6) (2022-10-20)


### Bug Fixes

* CVE-2022-32149 ([015c19a](https://github.com/open-feature/open-feature-operator/commit/015c19ac4455673902c365111816b021f893c485))

## [0.2.5](https://github.com/open-feature/open-feature-operator/compare/v0.2.4...v0.2.5) (2022-10-19)


### Features

* stop creation and mounting of flagd-config config map in case of kubernetes sync-provider ([#126](https://github.com/open-feature/open-feature-operator/issues/126)) ([a1d9fe2](https://github.com/open-feature/open-feature-operator/commit/a1d9fe276a37259d01e6ed6239c0ebcd3a1e6611))

## [0.2.4](https://github.com/open-feature/open-feature-operator/compare/v0.2.3...v0.2.4) (2022-10-18)


### Bug Fixes

* build and push to docker registry with tag as current release ([#123](https://github.com/open-feature/open-feature-operator/issues/123)) ([d4abda1](https://github.com/open-feature/open-feature-operator/commit/d4abda119e4a7c2dab7a2e0d335d44b1df07ec62))

## [0.2.3](https://github.com/open-feature/open-feature-operator/compare/v0.2.2...v0.2.3) (2022-10-18)


### Bug Fixes

* build and push to docker registry on tag creation ([#121](https://github.com/open-feature/open-feature-operator/issues/121)) ([27c6f9c](https://github.com/open-feature/open-feature-operator/commit/27c6f9cbc298fb8bf578464e4c3f9f07402b87ab))

## [0.2.2](https://github.com/open-feature/open-feature-operator/compare/v0.2.1...v0.2.2) (2022-10-14)


### Bug Fixes

* bump flagd version to include change detection fix ([421cab6](https://github.com/open-feature/open-feature-operator/commit/421cab651f6ebe2ece1380fda7dc24d92838d6b5))

## [0.2.1](https://github.com/open-feature/open-feature-operator/compare/v0.2.0...v0.2.1) (2022-10-13)


### Features

* metrics ([#111](https://github.com/open-feature/open-feature-operator/issues/111)) ([6016669](https://github.com/open-feature/open-feature-operator/commit/6016669ec46984d127951ee5d0ff02e7685f4d80))
* pr github action workflow ([#96](https://github.com/open-feature/open-feature-operator/issues/96)) ([a719f8a](https://github.com/open-feature/open-feature-operator/commit/a719f8a33abc9b9599987314282cc4e7ac202d67))


### Bug Fixes

* include assets in release ([#109](https://github.com/open-feature/open-feature-operator/issues/109)) ([b835abb](https://github.com/open-feature/open-feature-operator/commit/b835abb48ae8ca3c9c63abd51ae5614a4068c003))

## [0.2.0](https://github.com/open-feature/open-feature-operator/compare/v0.1.1...v0.2.0) (2022-10-10)


### ⚠ BREAKING CHANGES

* bump flagd version to 0.2.0 (connect refactor) (#97)

### Features

* bump flagd version to 0.2.0 (connect refactor) ([#97](https://github.com/open-feature/open-feature-operator/issues/97)) ([8118b9f](https://github.com/open-feature/open-feature-operator/commit/8118b9fcbaf0d3c66d6869369add645e388989de))


### Bug Fixes

* upgrade dependencies with vulnerabilities ([#90](https://github.com/open-feature/open-feature-operator/issues/90)) ([58cdd4e](https://github.com/open-feature/open-feature-operator/commit/58cdd4ee7c6989e44258bad3e9ed75a3bb465cae))
