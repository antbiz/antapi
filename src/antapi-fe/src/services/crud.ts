// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/**
 * 通用的 CRUD api
 */

/** 查询列表 */
export async function getMany(
  params: {
    collectionName: string,
    [key: string]: any
  },
) {
  const collectionName = params.collectionName;
  delete params['collectionName'];
  return request<{ data: object[] }>(`/api/biz/${collectionName}`, {
    method: 'GET',
    params: {
      ...params,
    },
  });
}

/** 查询详情 */
export async function getOne(
  collectionName: string,
  id: string,
  options?: { [key: string]: any }
) {
  return request<{ data: object }>(`/api/biz/${collectionName}/${id}`, {
    method: 'GET',
    ...(options || {}),
  });
}

/** 更新 */
export async function updateOne(
  collectionName: string,
  id: string,
  options?: { [key: string]: any }
) {
  return request<{ data: object }>(`/api/biz/${collectionName}/${id}`, {
    method: 'PUT',
    ...(options || {}),
  });
}

/** 新建 */
export async function createOne(
  collectionName: string,
  options?: { [key: string]: any }
) {
  return request<{ data: object }>(`/api/biz/${collectionName}`, {
    method: 'POST',
    ...(options || {}),
  });
}

/** 删除 */
export async function deleteOne(
  collectionName: string,
  id: string,
  options?: { [key: string]: any }
) {
  return request(`/api/biz/${collectionName}/${id}`, {
    method: 'DELETE',
    ...(options || {}),
  });
}
