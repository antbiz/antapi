// @ts-ignore
/* eslint-disable */

declare namespace API {
  type CurrentUser = {
    name?: string;
    avatar?: string;
    userid?: string;
    email?: string;
    signature?: string;
    title?: string;
    group?: string;
    tags?: { key?: string; label?: string }[];
    notifyCount?: number;
    unreadCount?: number;
    country?: string;
    access?: string;
    geographic?: {
      province?: { label?: string; key?: string };
      city?: { label?: string; key?: string };
    };
    address?: string;
    phone?: string;
  };

  type LoginResult = {
    status?: string;
    type?: string;
    currentAuthority?: string;
  };

  type PageParams = {
    current?: number;
    pageSize?: number;
  };

  type RuleListItem = {
    key?: number;
    disabled?: boolean;
    href?: string;
    avatar?: string;
    name?: string;
    owner?: string;
    desc?: string;
    callNo?: number;
    status?: number;
    updatedAt?: string;
    createdAt?: string;
    progress?: number;
  };

  type RuleList = {
    data?: RuleListItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type FakeCaptcha = {
    code?: number;
    status?: string;
  };

  type LoginParams = {
    username?: string;
    password?: string;
    autoLogin?: boolean;
    type?: string;
  };

  type ErrorResponse = {
    /** 业务约定的错误码 */
    errorCode: string;
    /** 业务上的错误信息 */
    errorMessage?: string;
    /** 业务上的请求是否成功 */
    success?: boolean;
  };

  type NoticeIconList = {
    data?: NoticeIconItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type NoticeIconItemType = 'notification' | 'message' | 'event';

  type NoticeIconItem = {
    id?: string;
    extra?: string;
    key?: string;
    read?: boolean;
    avatar?: string;
    title?: string;
    status?: string;
    datetime?: string;
    description?: string;
    type?: NoticeIconItemType;
  };

  /**
   * 模型字段描述
   */
  type SchemaField = {
    // 32 位 Id，需要手动生成
    id: string;
    // 字段类型
    type: SchemaFieldType;
    // 展示标题
    displayName: string;
    // 在数据库中的字段名
    name: string;
    // 字段顺序
    order: number;
    // 字段描述
    description: string;
    // 是否隐藏
    isHidden: boolean;
    // 是否必需字段
    isRequired: boolean;
    // 排序字段
    isOrderField: boolean;
    orderDirection: 'asc' | 'desc';
    // 是否为系统内置字段
    isSystem: boolean;
    // 是否唯一
    isUnique: boolean;
    // 在 API 返回结果中隐藏
    isHiddenInApi: boolean;
    // 是否加密
    isEncrypted: boolean;
    // 默认值
    defaultValue: any;
    // 最小长度/值
    min: number;
    // 最大长度/值
    max: number;
    // 校验
    validator: string;
    // 样式属性
    style: {};
    // 连接字段
    connectField: string;
    // 连接模型 Id
    connectResource: string;
    // 关联多个
    connectMany: boolean;
    // 枚举
    // 枚举元素的类型
    enumElementType: 'string' | 'number';
    // 所有枚举元素
    enumElements: { label: string; value: string }[];
    // 允许多个值
    isMultiple: boolean;
    // 图片、文件存储链接的形式，fileId 或 https 形式，默认为 true，
    resourceLinkType: 'fileId' | 'https';
    // 时间存储格式
    dateFormatType: 'timestamp-ms' | 'timestamp-s' | 'date' | 'string';
    // 多媒体类型
    mediaType: 'video' | 'music';
  };

  /**
   * 模型描述
   */
  type Schema = {
    id: string;
    createdAt: number;
    updatedAt: number;
    displayName: string;
    collectionName: string;
    projectId: string;
    // 在多个项目之间实现共享
    projectIds: string[];
    fields: SchemaField[];
    searchFields: SchemaField[];
    description: string;
    // 文档创建时间字段名
    docCreatedAtField: string;
    // 文件更新数据字段名
    docUpdatedAtField: string;
  };

  type Project = {
    id: string;
    createdAt: number;
    updatedAt: number;
    name: string;
    customId: string;
    description: string;
    // 项目封面图
    cover?: string;
    // 是否开启 Api 访问
    enableApiAccess: boolean;
    // api 访问路径
    apiAccessPath: string;
    // 可读集合
    readableCollections: string[];
    // 可修改的集合
    modifiableCollections: string[];
    // 可删除的集合
    deletableCollections: string[];
    keepApiPath: boolean;
  };

  type User = {
    id: string;
    username: string;
    // 创建时间
    createdAt: number;
    // 用户角色
    roles: UserRole[];
    // uuid
    uuid: string;
    // 是否为 root 用户
    root?: boolean;
  };

  type UserRole = {
    id: string;
    // 角色名
    roleName: string;
    // 角色描述
    description: string;
    // 角色绑定的权限描述
    permissions: Permission[];
    type: string | 'system';
  };
}
