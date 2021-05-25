// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 修改个人信息 GET /api/user/info */
export async function updateInfo(data?: { [key: string]: any }) {
  return request<API.CurrentUser>('/api/user/info', {
    method: 'POST',
    data: data,
  });
}

/** 修改个人密码 GET /api/user/update_password */
export async function updatePassword(data?: { [key: string]: any }) {
  return request<API.CurrentUser>('/api/user/update_password', {
    method: 'POST',
    data: data,
  });
}
