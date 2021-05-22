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

  type Project = {
    id: string;
    createdAt: string;
    updatedAt: string;
    title: string;
    name: string;
    description: string;
    cover?: string;
  };

  /**
   * 模型描述
   */
  type Schema = {
    id: string;
    createdAt: string;
    updatedAt: string;
    projectName: string;
    title: string;
    collectionName: string;
    description: string;
    displayType?: string;
    showDescIcon?: boolean;
    type?: string;
    properties?: object;
  };

  /**
   * 权限
   */
  type Permission = {
    id: string;
    createdAt: string;
    updatedAt: string;
    title: string;
    projectName: string;
    collectionName: string;
    createLevel: int;
    readLevel: int;
    updateLevel: int;
    deleteLevel: int;
  }
}
