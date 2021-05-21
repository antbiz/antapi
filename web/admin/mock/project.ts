// eslint-disable-next-line import/no-extraneous-dependencies
import { Request, Response } from 'express';
import { parse } from 'url';

// mock ProjectListDataSource
const genList = (current: number, pageSize: number) => {
  const projectListDataSource: API.Project[] = [];

  for (let i = 1; i < pageSize; i += 1) {
    const index = (current - 1) * 10 + i;
    projectListDataSource.push({
      id: index.toString(),
      createdAt: '2021-03-28',
      updatedAt: '2021-03-28',
      title: `使用示例 ${index}`,
      cover: ['https://gw.alipayobjects.com/zos/rmsportal/eeHMaZBwmTvLdIwMfBpg.png', ''][i % 2],
      name: `example${index}`,
      description: 'CMS使用示例',
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
  'GET /api/biz/project': getProjects,
  'GET /api/biz/project/1': {
    data: {
      id: "1",
      createdAt: '2021-03-28',
      updatedAt: '2021-03-28',
      title: `使用示例1`,
      cover: 'https://gw.alipayobjects.com/zos/rmsportal/eeHMaZBwmTvLdIwMfBpg.png',
      name: `example1`,
      description: 'CMS使用示例',
    },
  },
  'GET /api/biz/project/example1': {
    data: {
      id: "1",
      createdAt: '2021-03-28',
      updatedAt: '2021-03-28',
      title: `使用示例1`,
      cover: 'https://gw.alipayobjects.com/zos/rmsportal/eeHMaZBwmTvLdIwMfBpg.png',
      name: `example1`,
      description: 'CMS使用示例',
    },
  },
  'GET /api/biz/permission': {
    data: [
      {
        id: '624748504',
        createdAt: '2020-05-26T09:42:56Z',
        updatedAt: '2020-05-26T08:19:22Z',
        title: '测试A',
        projectName: 'example1',
        collectionName: 'test_a',
        createLevel: 0,
        readLevel: 0,
        updateLevel: 0,
        deleteLevel: 0,
      },
      {
        id: '624691229',
        createdAt: '2020-05-26T08:19:22Z',
        updatedAt: '2020-05-26T08:19:22Z',
        title: '测试B',
        projectName: 'example1',
        collectionName: 'test_b',
        createLevel: 0,
        readLevel: 0,
        updateLevel: 0,
        deleteLevel: 0,
      },
    ]
  }
};
