

下面是 OpenList API 文档中各接口分类及其作用的详细说明：

---

### **Authentication（认证）**

- **User login** `POST` — 用户通过用户名和密码进行标准登录，获取 JWT 令牌。
- **User login with pre-hashed password** `POST` — 使用预哈希密码登录，适用于密码已在客户端完成哈希处理的场景，避免明文传输密码。
- **LDAP login** `POST` — 通过 LDAP（轻量级目录访问协议）服务器进行身份验证登录，适用于企业内部统一认证体系。
- **User logout** `GET` — 用户登出，使当前会话或令牌失效。
- **Generate 2FA secret** `POST` — 为当前用户生成双因素认证（2FA）的密钥（通常是一个 TOTP 密钥和二维码），用于绑定验证器应用。
- **Verify and enable 2FA** `POST` — 验证用户输入的 2FA 验证码是否正确，验证通过后正式启用双因素认证。
- **SSO login redirect** `GET` — 将用户重定向到 SSO（单点登录）提供商的认证页面，发起 SSO 登录流程。
- **SSO callback handler** `GET` — 处理 SSO 提供商认证完成后的回调，接收授权码或令牌，完成 SSO 登录。
- **Begin WebAuthn login** `GET` — 发起 WebAuthn（无密码认证）登录流程，返回认证挑战（challenge），供浏览器调用身份验证器。
- **Finish WebAuthn login** `POST` — 完成WebAuthn 登录，提交身份验证器的签名响应，服务端验证后签发令牌。
- **Begin WebAuthn registration** `GET` — 发起 WebAuthn 凭据注册流程，返回注册挑战，供用户绑定新的安全密钥或设备。
- **Finish WebAuthn registration** `POST` — 完成 WebAuthn 凭据注册，提交身份验证器的注册数据，服务端保存凭据。
- **Delete WebAuthn credential** `POST` — 删除指定的 WebAuthn 凭据，解绑某个安全密钥或设备。
- **Get WebAuthn credentials** `GET` — 获取当前用户已注册的所有 WebAuthn 凭据列表。

---

### **User（用户自身操作）**

- **Get current user info** — 获取当前登录用户的个人信息（用户名、权限、偏好等）。
- **Update current user** — 更新当前用户的个人资料，如昵称、密码、偏好设置等。
- **List my SSH public keys** — 列出当前用户已添加的所有 SSH 公钥。
- **Add SSH public key** — 为当前用户添加一个新的 SSH 公钥，可用于 SSH 方式访问存储后端。
- **Delete SSH public key** — 删除当前用户的某个 SSH 公钥。

---

### **Admin — 用户管理**

- **List all users (Admin)** `GET` — 管理员获取系统中所有用户的列表（支持分页）。
- **Get user by ID (Admin)** `GET` — 管理员根据用户 ID 获取单个用户的详细信息。
- **Create new user (Admin)** `POST` — 管理员创建新用户账号，指定用户名、密码、权限等。
- **Update user (Admin)** `POST` — 管理员更新指定用户的信息，如修改权限、禁用账号等。
- **Delete user (Admin)** `POST` — 管理员删除指定用户账号。
- **Cancel user 2FA (Admin)** `POST` — 管理员取消/重置指定用户的双因素认证（在用户丢失验证器时使用）。
- **Clear user cache (Admin)** `POST` — 管理员清除指定用户的缓存数据，强制刷新。
- **List user SSH keys (Admin)** `GET` — 管理员列出指定用户的所有 SSH 公钥。
- **Delete user SSH key (Admin)** `POST` — 管理员删除指定用户的某个 SSH 公钥。

---

### **Admin — 存储管理**

- **List all storages (Admin)** `GET` — 管理员获取所有已挂载存储的列表。
- **Get storage by ID (Admin)** `GET` — 管理员根据存储 ID 获取单个存储的详细配置信息。
- **Create storage (Admin)** `POST` — 管理员创建新的存储挂载，将外部存储（如本地磁盘、对象存储、网盘等）接入系统。
- **Update storage (Admin)** `POST` — 管理员更新指定存储的配置，如修改挂载路径、驱动参数等。
- **Delete storage (Admin)** `POST` — 管理员删除指定的存储挂载。
- **Enable storage (Admin)** `POST` — 管理员启用指定的存储挂载，使其可用。
- **Disable storage (Admin)** `POST` — 管理员禁用指定的存储挂载，暂停其使用但不删除配置。
- **Reload all storages (Admin)** `POST` — 管理员重新加载所有存储配置，使配置变更生效。

---

### **Admin — 驱动管理**

- **List all drivers (Admin)** `GET` — 获取所有可用的存储驱动及其配置模板（当前页面展示的接口）。
- **Get driver names (Admin)** `GET` — 仅获取所有驱动的名称列表，比完整列表更轻量。
- **Get driver info (Admin)** `GET` — 获取指定驱动的详细信息，包括其支持的配置项和参数说明。

---

### **Admin — 系统设置**

- **List all settings (Admin)** `GET` — 管理员获取系统所有配置项的列表。
- **Get setting by key (Admin)** `GET` — 管理员根据配置键名获取单个配置项的值。
- **Save settings (Admin)** `POST` — 管理员保存/更新系统配置项。
- **Delete setting (Admin)** `POST` — 管理员删除指定的配置项（恢复为默认值）。
- **Reset API token (Admin)** `POST` — 管理员重置系统的 API 令牌，旧令牌即刻失效。

---

### **Admin — 元数据管理**

- **List all metas (Admin)** `GET` — 管理员获取所有元数据配置的列表（元数据用于为特定目录/文件自定义显示或行为）。
- **Get meta by ID (Admin)** `GET` — 管理员根据 ID 获取单个元数据配置的详情。
- **Create meta (Admin)** `POST` — 管理员创建新的元数据配置，如为某个路径设置密码、说明、排序规则等。
- **Update meta (Admin)** `POST` — 管理员更新指定的元数据配置。
- **Delete meta (Admin)** `POST` — 管理员删除指定的元数据配置。

---

### **Admin — 搜索索引管理**

- **Build search index (Admin)** `POST` — 管理员触发全量搜索索引构建，从头扫描所有文件建立索引。
- **Update search index (Admin)** `POST` — 管理员触发增量搜索索引更新，只处理变更部分。
- **Stop indexing (Admin)** `POST` — 管理员停止正在进行的索引构建/更新任务。
- **Clear search index (Admin)** `POST` — 管理员清除所有搜索索引数据。
- **Get indexing progress (Admin)** `GET` — 管理员获取当前索引任务的进度信息。

---

### **File System（文件系统操作）**

- **List directory contents** — 列出指定目录下的文件和子目录。
- **Get file or directory info** — 获取指定文件或目录的详细信息（大小、修改时间等）。
- **Search files and directories** — 按关键词搜索文件和目录。
- **Get directory tree** — 获取目录的树形结构，以层级方式展示文件组织。
- **Get additional file operations** — 获取文件支持的其他操作（如直链、预览等）。
- **Create directory** — 创建新目录。
- **Rename file or directory** — 重命名指定的文件或目录。
- **Batch rename files** — 批量重命名多个文件。
- **Regex-based rename** — 使用正则表达式匹配并重命名文件。
- **Move files or directories** — 移动文件或目录到目标路径。
- **Recursive move** — 递归移动，将目录及其所有内容一起移动。
- **Copy files or directories** — 复制文件或目录到目标路径。
- **Remove files or directories** — 删除指定的文件或目录。
- **Remove empty directories** — 删除空目录（不删除含文件的目录）。
- **Upload file (stream)** — 以流式方式上传文件，适合大文件传输。
- **Upload file (form)** — 以表单方式上传文件，适合小文件或浏览器端上传。
- **Add offline download task** — 添加离线下载任务，由服务器端下载指定 URL 的文件到存储中。
- **Decompress archive** — 解压缩归档文件（如 zip、tar.gz 等）。
- **Get archive metadata** — 获取归档文件的元数据信息，不解压。
- **List archive contents** — 列出归档文件内的文件列表，不解压。

---

### **Public（公开接口）**

- **Get public settings** — 获取公开的系统设置信息（无需认证），如站点名称、公告等。
- **Get available offline download tools** — 获取系统可用的离线下载工具列表（如 aria2 等）。
- **Get supported archive extensions** — 获取系统支持的归档文件扩展名列表。

---

### **Sharing（分享管理）**

- **List all shares** — 列出所有文件分享记录。
- **Get share by ID** — 根据分享 ID 获取单个分享的详细信息。
- **Create file share** — 创建新的文件/目录分享链接。
- **Update share** — 更新分享的配置（如修改密码、过期时间等）。
- **Delete share** — 删除指定的分享。
- **Enable share** — 启用指定的分享，使其可被访问。
- **Disable share** — 禁用指定的分享，暂停访问但不删除记录。

---

### **TS 版本接口**

- **文件操作接口** — TypeScript 版本的文件操作接口封装。
- **用户操作接口** — TypeScript 版本的用户操作接口封装。
- **挂载管理接口** — TypeScript 版本的存储挂载管理接口封装。

---

### **密钥站接口**

文档中未展开详细说明，推测为与密钥/许可证管理相关的专用接口。

---

### **Schemas（数据模型）**

这些不是接口，而是 API 请求/响应中使用的数据结构定义：

- **ApiResponse** — 通用 API 响应结构（包含 code、message、data）。
- **ErrorResponse** — 错误响应结构。
- **PageReq** — 分页请求参数（页码、每页数量）。
- **Pagination** — 分页元数据（总条数、总页数等）。
- **User** — 用户数据模型。
- **LoginRequest / LoginResponse** — 登录请求/响应模型。
- **UserResponse / UsersListResponse** — 用户信息/用户列表响应模型。
- **FsObject** — 文件系统对象模型（文件或目录的属性）。
- **FsListRequest / FsListResponse** — 目录列表请求/响应模型。
- **FsGetRequest / FsGetResponse** — 文件信息获取请求/响应模型。
- **FsMkdirRequest** — 创建目录请求模型。
- **FsRenameRequest** — 重命名请求模型。
- **FsMoveCopyRequest** — 移动/复制请求模型。
- **FsRemoveRequest** — 删除请求模型。
- **StorageDetails / Storage** — 存储配置详情/摘要模型。
- **DriverInfo** — 驱动信息模型（包含驱动名称和配置模板）。
- **StorageResponse / StoragesListResponse** — 存储信息/存储列表响应模型。