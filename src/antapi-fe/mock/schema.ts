import { Request, Response } from 'express';

export default {
  'GET /api/project/0/schema': {
    data: [
      {
        id: '0',
        displayName: '图片',
        collectionName: 'img',
        docCreateTimeField: 'createdAt',
        docUpdateTimeField: 'updatedAt',
        fields: [
          {
            displayName: '枚举',
            enumElementType: 'string',
            enumElements: [
              {
                label: '自然',
                value: 'nature',
              },
              {
                label: '人物',
                value: 'person',
              },
            ],
            id: 'rbalbfgld6uu8ga8l565h3kcjes0yufa',
            name: 'enum',
            order: 0,
            type: 'Enum',
          },
          {
            displayName: '图片',
            id: 'b4bn57eqmwtndnm1u02aqgxo7epxohtx',
            isMultiple: true,
            name: 'img',
            order: 1,
            resourceLinkType: 'fileId',
            type: 'Image',
          },
          {
            dateFormatType: 'timestamp-ms',
            description: '系统字段，请勿随意修改',
            displayName: '创建时间',
            id: 'createdAt',
            isSystem: true,
            name: 'createdAt',
            type: 'DateTime',
          },
          {
            dateFormatType: 'timestamp-ms',
            description: '系统字段，请勿随意修改',
            displayName: '修改时间',
            id: 'updatedAt',
            isSystem: true,
            name: 'updatedAt',
            type: 'DateTime',
          },
          {
            connectField: 'aasd',
            connectResource: '3b020ca3601b6354026caffd2bd94f1a',
            displayName: '视频',
            id: 'ac2h3uaifbrbenla7hazkd4k7zxbqw9a',
            name: 'aasd',
            order: 4,
            type: 'Connect',
          },
          {
            displayName: '1',
            id: 'cr9d8vk4xyrsuicbty6cspa6dbiap4g2',
            name: '1',
            order: 5,
            type: 'Object',
          },
        ],
        createdAt: 1611624969143,
        updatedAt: 1614672363463,
        projectId: '0',
      },
      {
        id: '1',
        collectionName: 'aasd',
        displayName: '视频',
        docCreateTimeField: 'createdAt',
        docUpdateTimeField: 'updatedAt',
        fields: [
          {
            dateFormatType: 'timestamp-ms',
            description: '系统字段，请勿随意修改',
            displayName: '创建时间',
            id: 'createdAt',
            isSystem: true,
            name: 'createdAt',
            type: 'DateTime',
          },
          {
            dateFormatType: 'timestamp-ms',
            description: '系统字段，请勿随意修改',
            displayName: '修改时间',
            id: 'updatedAt',
            isSystem: true,
            name: 'updatedAt',
            type: 'DateTime',
          },
          {
            displayName: '视频',
            id: 'hh52ncwuj5unlmok85yndcve1mx7dw85',
            isMultiple: true,
            isRequired: true,
            mediaType: 'video',
            name: 'aasd',
            order: 2,
            type: 'Media',
          },
        ],
        createdAt: 1612407636000,
        updatedAt: 1612407679924,
        projectId: '0',
      },
      {
        id: '2',
        collectionName: 'activities',
        description: '测试用',
        displayName: '活动',
        docCreateTimeField: 'createdAt',
        docUpdateTimeField: 'updatedAt',
        fields: [
          {
            dateFormatType: 'timestamp-ms',
            description: '系统字段，请勿随意修改',
            displayName: '创建时间',
            id: 'createdAt',
            isSystem: true,
            name: 'createdAt',
            type: 'DateTime',
          },
          {
            dateFormatType: 'timestamp-ms',
            description: '系统字段，请勿随意修改',
            displayName: '修改时间',
            id: 'updatedAt',
            isSystem: true,
            name: 'updatedAt',
            type: 'DateTime',
          },
          {
            defaultValue: '体验活动',
            description: '活动名称',
            displayName: '活动名称',
            id: 'ebf5ft6pt5nkq4etdzmkqddo4g9pdh2h',
            isRequired: true,
            name: 'act_name',
            order: 2,
            type: 'String',
          },
          {
            displayName: '活动简介',
            id: 'zr8w9zhjre9fcg0p59t5ld5ufsmawte0',
            name: 'act_context',
            order: 3,
            type: 'MultiLineString',
          },
          {
            dateFormatType: 'string',
            displayName: '活动时间',
            id: 'd9o2hoeuxkze094xvm88nn9qqjiil4la',
            name: 'act_time',
            order: 4,
            type: 'DateTime',
          },
          {
            description:
              '可使用低码开发并发布的H5页面作为跳转中间页，或直接用单张图片作为跳转中间页',
            displayName: '跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '高级自定义H5页面（由低码开发）',
              },
              {
                label: '单图片跳转',
              },
            ],
            id: 'ju6sy7saub54o8in7t97yg9vm86fsoef',
            isRequired: true,
            name: 'nb',
            order: 5,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb',
            order: 6,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb2',
            order: 7,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb3',
            order: 8,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb4',
            order: 9,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb5',
            order: 10,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb6',
            order: 11,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb7',
            order: 12,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb8',
            order: 12,
            type: 'Enum',
          },
          {
            defaultValue: '“招聘猎头”应用-个人中心(user)',
            displayName: '高级跳转中间页',
            enumElementType: 'string',
            enumElements: [
              {
                label: '“招聘猎头”应用-候选列表(index)',
              },
              {
                label: '“招聘猎头”应用-岗位信息(graph)',
              },
              {
                label: '“招聘猎头”应用-个人中心(user)',
              },
            ],
            id: 'lgjcfigpekax80l3arcn6a3tqyvs41lg',
            isRequired: true,
            name: 'nbnb9',
            order: 13,
            type: 'Enum',
          },
        ],
        createdAt: 1614239812022,
        updatedAt: 1614673719350,
        projectId: '0',
      },
    ],
  },
};
