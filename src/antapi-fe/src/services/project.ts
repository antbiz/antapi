// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取项目列表 GET /api/project */
export async function getProjects(
  params?: {
    // query
    /** 当前的页码 */
    current?: number;
    /** 页面的容量 */
    pageSize?: number;
  },
  options?: { [key: string]: any },
) {
  return request<{ data: API.Project[] }>('/api/project', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获取项目详情  GET /api/project */
export async function getProject(projectId: string, options?: { [key: string]: any }) {
  return request<{ data: API.Project }>(`/api/project/${projectId}`, {
    method: 'GET',
    ...(options || {}),
  });
}

/** 新建项目 PUT /api/project */
export async function updateProject(options?: { [key: string]: any }) {
  return request<{ data: API.Project }>('/api/project', {
    method: 'PUT',
    ...(options || {}),
  });
}

/** 新建项目 POST /api/project */
export async function createProject(options?: { [key: string]: any }) {
  return request<{ data: API.Project }>('/api/project', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 删除项目 DELETE /api/project */
export async function deleteProject(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/project', {
    method: 'DELETE',
    ...(options || {}),
  });
}
