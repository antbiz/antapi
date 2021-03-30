// eslint-disable-next-line import/no-extraneous-dependencies
import { Request, Response } from 'express';
import { parse } from 'url';

// mock ProjectListDataSource
const genList = (current: number, pageSize: number) => {
  const projectListDataSource: API.Project[] = [];

  for (let i = 0; i < pageSize; i += 1) {
    const index = (current - 1) * 10 + i;
    projectListDataSource.push({
      id: index.toString(),
      cover: ['https://gw.alipayobjects.com/zos/rmsportal/eeHMaZBwmTvLdIwMfBpg.png', ''][i % 2],
      name: `使用示例 ${index}`,
      customId: `cmstest${index}`,
      createdAt: 1612438962519,
      updatedAt: 1612438962519,
      description: 'CMS使用示例',
      enableApiAccess: true,
      apiAccessPath: 'api1',
      keepApiPath: false,
      deletableCollections: [],
      modifiableCollections: [],
      readableCollections: ['news'],
    });
  }
  return projectListDataSource;
};

function getProjects(req: Request, res: Response, u: string) {
  const projectListDataSource = genList(1, 10);
  let realUrl = u;
  if (!realUrl || Object.prototype.toString.call(realUrl) !== '[object String]') {
    realUrl = req.url;
  }
  const { current = 1, pageSize = 10 } = req.query;
  const params = (parse(realUrl, true).query as unknown) as API.PageParams &
    API.Project[] & {
      sorter: any;
      filter: any;
    };
  const result = {
    data: projectListDataSource,
    total: projectListDataSource.length,
    success: true,
    pageSize,
    current: parseInt(`${params.current}`, 10) || 1,
  };

  return res.json(result);
}

export default {
  'GET /api/project': getProjects,
  'GET /api/project/0': {
    data: {
      id: '79550af2600e8de600da8ad13724cd98',
      customId: 'sms',
      name: '短信测试',
      createdAt: 1611566566279,
      enableApiAccess: true,
    },
  },
};
