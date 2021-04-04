// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取内容模型列表 GET `/api/biz/schema?projectName=${projectName}` */
export async function getSchemas(
  params: {
    projectName: string
  }
) {
  return request<{ data: API.Schema[] }>(`/api/biz/schema?projectName[eq]${params.projectName}`, {
    method: 'GET',
  });
}

/** 获取内容模型详情  GET `/api/biz/schema/${schemaId}` */
export async function getSchema(schemaId: string) {
  return request<{ data: API.Schema }>(`/api/biz/schema/${schemaId}`, {
    method: 'GET',
  });
}

/** 新建内容模型 PUT `/api/biz/schema/${schemaId}`*/
export async function updateSchema(
  schemaId: string,
  options?: { [key: string]: any },
) {
  return request<{ data: API.Schema }>(`/api/biz/schema/${schemaId}`, {
    method: 'PUT',
    ...(options || {}),
  });
}

/** 新建内容模型 POST `/api/biz/schema`*/
export async function createSchema(
  options?: { [key: string]: any },
) {
  return request<{ data: API.Schema }>(`/api/biz/schema`, {
    method: 'POST',
    ...(options || {}),
  });
}

/** 删除内容模型 DELETE `/api/biz/schema/${schemaId}`*/
export async function deleteSchema(
  schemaId: string,
  options?: { [key: string]: any },
) {
  return request(`/api/biz/schema/${schemaId}`, {
    method: 'DELETE',
    ...(options || {}),
  });
}
