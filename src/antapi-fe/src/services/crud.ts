// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/**
 * 通用的 CRUD api
 */

/** 查询列表 */
export async function getMany(
  schemaName: string,
  params?: { [key: string]: any },
) {
  return request<{ data: object[] }>(`/api/admin/${schemaName}`, {
    method: 'GET',
    params: {
      ...params,
    },
  });
}

/** 查询详情 */
export async function getOne(
  schemaName: string,
  id: string,
  options?: { [key: string]: any }
) {
  return request<{ data: object }>(`/api/admin/${schemaName}/${id}`, {
    method: 'GET',
    ...(options || {}),
  });
}

/** 更新 */
export async function updateOne(
  schemaName: string,
  id: string,
  options?: { [key: string]: any }
) {
  return request<{ data: object }>(`/api/admin/${schemaName}/${id}`, {
    method: 'PUT',
    ...(options || {}),
  });
}

/** 新建 */
export async function createOne(options?: { [key: string]: any }) {
  return request<{ data: object }>('/api/admin/project', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 删除 */
export async function deleteOne(
  schemaName: string,
  id: string,
  options?: { [key: string]: any }
) {
  return request(`/api/admin/${schemaName}/${id}`, {
    method: 'DELETE',
    ...(options || {}),
  });
}
