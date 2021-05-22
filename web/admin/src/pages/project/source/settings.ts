export const globalSettings={{
  type: 'object',
  properties: {
    title: {
      title: '标题',
      description: '模型展示名称',
      type: 'string',
      required: true,
    },
    name: {
      title: '模型名称/英文',
      description: '',
      type: 'string',
      required: true,
    },
    projectName: {
      title: '项目名称/英文',
      type: 'string',
      required: true,
    },
    description: {
      title: '模型描述',
      type: 'string',
      format: "textarea",
    },
    column: {
      title: '整体布局',
      type: 'number',
      enum: [1, 2, 3],
      enumNames: ['一行一列', '一行二列', '一行三列'],
      props: {
        placeholder: '默认一行一列',
      },
    },
    labelWidth: {
      title: '标签宽度',
      type: 'number',
      widget: 'slider',
      max: 300,
      props: {
        hideNumber: true,
      },
    },
    displayType: {
      title: '标签展示模式',
      type: 'string',
      enum: ['row', 'column'],
      enumNames: ['同行', '单独一行'],
      widget: 'radio',
    },
  }
};

export const commonSettings = {
  _id: {
    type: 'string',
    hidden: true,
  },
  createdAt: {
    type: 'string',
    hidden: true,
  },
  updatedAt: {
    type: 'string',
    hidden: true,
  },
  createdBy: {
    type: 'string',
    hidden: true,
  },
  updatedBy: {
    type: 'string',
    hidden: true,
  },
  $id: {
    title: 'ID',
    description: '字段名称/英文',
    type: 'string',
    widget: 'idInput',
    required: true,
  },
  title: {
    title: '标题',
    type: 'string',
  },
  description: {
    title: '说明',
    type: 'string',
  },
  default: {
    title: '默认值',
    type: 'string',
  },
  required: {
    title: '必填',
    type: 'boolean',
  },
  placeholder: {
    title: '占位符',
    type: 'string',
  },
  bind: {
    title: 'Bind',
    type: 'string',
  },
  min: {
    title: '最小值',
    type: 'number',
  },
  max: {
    title: '最大值',
    type: 'number',
  },
  disabled: {
    title: '禁用',
    type: 'boolean',
  },
  readOnly: {
    title: '只读',
    type: 'boolean',
  },
  hidden: {
    title: '隐藏',
    type: 'boolean',
  },
  width: {
    title: '元素宽度',
    type: 'string',
    widget: 'percentSlider',
  },
  labelWidth: {
    title: '标签宽度',
    description: '默认值120',
    type: 'number',
    widget: 'slider',
    max: 400,
    props: {
      hideNumber: true,
    },
  },
}
