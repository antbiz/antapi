import { useParams } from 'umi';
import React, { useMemo, useState } from 'react';
import { FieldTypes } from '@/common';
import { Card, Space, Typography, Tooltip, Switch, Tag, Spin } from 'antd';
import { ExclamationCircleTwoTone, QuestionCircleTwoTone } from '@ant-design/icons';
import type { DropResult } from 'react-beautiful-dnd';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { getSchemaSystemFields } from '@/utils';

export interface FieldType {
  icon: React.ReactNode;
  name: string;
  type: string;
}

const { Paragraph, Text, Title } = Typography;

export const SchemaSystemField: React.FC<{ onFiledClick: Function; schema: Schema }> = ({
  onFiledClick,
  schema,
}) => {
  const [showSystemField, setShowSystemField] = useState(false);

  // 合并系统字段
  const systemFields = useMemo(() => getSchemaSystemFields(schema), [schema]);

  return (
    <div>
      <Paragraph>
        <Space>
          <Switch checked={showSystemField} onChange={(v) => setShowSystemField(v)} />
          <span>展示系统功能性字段</span>
          <Tooltip title="系统功能性字段为系统功能相关的字段，请谨慎操作">
            <QuestionCircleTwoTone />
          </Tooltip>
        </Space>
      </Paragraph>
      {showSystemField
        ? systemFields.map((field, index) => {
            const type = FieldTypes.find((_) => _.type === field.type);
            return (
              <Card
                hoverable
                key={index}
                className="schema-field-card system-field"
                onClick={() => onFiledClick(field)}
              >
                <Space style={{ flex: '1 1 auto' }}>
                  <div className="icon">{type?.icon}</div>
                  <div className="flex-column">
                    <Space align="center" style={{ marginBottom: '10px' }}>
                      <Tooltip title={field.displayName}>
                        <Title ellipsis level={4} style={{ marginBottom: 0 }}>
                          {field.displayName}
                        </Title>
                      </Tooltip>
                      <Text strong># {field.name}</Text>
                      <Tooltip title="系统字段，请勿随意修改">
                        <ExclamationCircleTwoTone style={{ fontSize: '16px' }} />
                      </Tooltip>
                    </Space>
                    <Space>
                      <Tag color="#9da6c7">{type?.name}</Tag>
                      <Tag color="#2575e6">系统字段</Tag>
                    </Space>
                  </div>
                </Space>
              </Card>
            );
          })
        : ''}
    </div>
  );
};

export const SchemaFieldRender: React.FC<{
  schema: Schema;
  onFiledClick: (filed: SchemaField) => void;
  actionRender: (field: SchemaField) => React.ReactNode;
}> = (props) => {
  const { schema, actionRender, onFiledClick } = props;
  const { projectId } = useParams<{projectId: string}>();
  console.log(projectId);

  const handleDragSort = async (result: DropResult) => {
    // 脱离列表，无效
    if (!result.destination) {
      return;
    }

    const source = result.source.index;
    const destination = result.destination.index;

    // 未改变顺序，无效
    if (source === destination) {
      return;
    }

    // 获取字段排序后的列表，系统字段不参与排序
    let resortedFields = schema?.fields
      .filter((_) => _ && !_.isSystem)
      .sort((prev, next) => prev.order - next.order);

    // 将被移动的字段移到对应的位置
    const moveField = resortedFields.splice(source, 1)?.[0];
    resortedFields.splice(destination, 0, moveField);

    // 重置 order 值，并添加系统字段
    resortedFields = resortedFields
      .map((field, index) => ({
        ...field,
        order: index,
      }))
      .concat(getSchemaSystemFields(schema));
  };

  return (
    <div>
      <SchemaSystemField {...props} />
      <Spin tip="加载中" spinning={loading || sortLoading}>
        <DragDropContext onDragEnd={handleDragSort}>
          <Droppable droppableId="droppable">
            {(droppableProvided) => (
              <div ref={droppableProvided.innerRef}>
                {schema?.fields
                  ?.filter((_) => _ && !_.isSystem)
                  .sort((prev, next) => prev.order - next.order)
                  .map((field, index) => {
                    const type = FieldTypes.find((_) => _.type === field.type);

                    return (
                      <Draggable key={field.id} draggableId={field.id} index={index}>
                        {(draggableProvided) => (
                          <div
                            className="schema-field-card"
                            ref={draggableProvided.innerRef}
                            {...draggableProvided.draggableProps}
                            {...draggableProvided.dragHandleProps}
                          >
                            <Card hoverable key={index} onClick={() => onFiledClick(field)}>
                              <Space style={{ flex: '1 1 auto' }}>
                                <div className="icon">{type?.icon}</div>
                                <div className="flex-column">
                                  <Space align="center" style={{ marginBottom: '10px' }}>
                                    <Tooltip title={field.displayName}>
                                      <Title ellipsis level={4} style={{ marginBottom: 0 }}>
                                        {field.displayName}
                                      </Title>
                                    </Tooltip>
                                    <Text strong># {field.name}</Text>
                                    {field.description && (
                                      <Tooltip title={field.description}>
                                        <ExclamationCircleTwoTone style={{ fontSize: '16px' }} />
                                      </Tooltip>
                                    )}
                                  </Space>
                                  <Space>
                                    <Tag>{type?.name}</Tag>
                                  </Space>
                                </div>
                              </Space>
                              {actionRender(field)}
                            </Card>
                          </div>
                        )}
                      </Draggable>
                    );
                  })}
                {droppableProvided.placeholder}
              </div>
            )}
          </Droppable>
        </DragDropContext>
      </Spin>
    </div>
  );
};
