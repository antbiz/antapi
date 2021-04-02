import { history, matchPath } from 'umi';

/* eslint no-useless-escape:0 import/prefer-default-export:0 */
const reg = /(((^https?:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+(?::\d+)?|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)$/;

export const isUrl = (path: string): boolean => reg.test(path);

/**
 * 从 url 中获取项目 id
 */
export const getProjectId = () => {
  const match = matchPath<{ projectId?: string }>(history.location.pathname, {
    path: '/project/:projectId/*',
    exact: true,
  });

  // 项目 Id
  const { projectId = '' } = match?.params || {};

  return projectId;
};
