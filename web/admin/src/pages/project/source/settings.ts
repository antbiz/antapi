export const globalSettings={
  type: 'object',
  properties: {
    projectName: {
      title: '项目ID/英文',
      type: 'string',
      disabled: true,
      required: true,
    },
    name: {
      title: '模型ID/英文',
      type: 'string',
      required: true,
    },
    displayName: {
      title: '模型名称',
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
  $id: {
    title: 'ID',
    description: '字段ID/英文',
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
  unique: {
    title: '唯一',
    type: 'boolean',
  },
  private: {
    title: '字段在api中不可见',
    type: 'boolean'
  },
  indexField: {
    title: '索引',
    type: 'boolean'
  },
  validator: {
    title: '校验器',
    type: 'string'
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
