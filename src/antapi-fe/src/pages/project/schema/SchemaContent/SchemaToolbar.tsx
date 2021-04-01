import { useParams } from 'umi';
import React, { useState, useEffect } from 'react';
import { Modal, Space, Checkbox, Typography, Tooltip } from 'antd';
import { EditTwoTone, DeleteTwoTone } from '@ant-design/icons';

const iconStyle = {
  fontSize: '16px',
};

/**
 * 删除模型弹窗
 */
export const DeleteSchemaModal: React.FC<{
  visible: boolean;
  currentSchema: API.Schema;
  onClose: () => void;
}> = ({ visible, currentSchema, onClose }) => {
  const { projectId } = useParams<any>();
  const [deleteCollection, setDeleteCollection] = useState(false);

  useEffect(() => {
    setDeleteCollection(false);
  }, [visible]);

  return (
    <Modal
      centered
      title="删除内容模型"
      visible={visible}
      onCancel={() => onClose()}
      onOk={async () => {
        console.log(projectId)
      }}
    >
      <Space direction="vertical">
        <Typography.Text>
          确认删【{currentSchema?.displayName} ({currentSchema?.collectionName})】内容模型？
        </Typography.Text>
        <Checkbox
          checked={deleteCollection}
          onChange={(e) => setDeleteCollection(e.target.checked)}
        >
          同时删除数据表（警告：删除后数据无法找回）
        </Checkbox>
      </Space>
    </Modal>
  );
};

const SchemaEdit: React.FC<{currentSchema: API.Schema}> = ({currentSchema}) => {
  // 删除模型
  const [deleteSchemaVisible, setDeleteSchemaVisible] = useState(false);

  return (
    <>
      <Space size="middle">
        {/* 编辑模型 */}
        <Tooltip title="编辑模型">
          <EditTwoTone
            style={iconStyle}
            onClick={() => {
            }}
          />
        </Tooltip>
        {/* 删除模型 */}
        <Tooltip title="删除模型">
          <DeleteTwoTone style={iconStyle} onClick={() => setDeleteSchemaVisible(true)} />
        </Tooltip>
      </Space>
      <DeleteSchemaModal
        visible={deleteSchemaVisible}
        currentSchema={currentSchema}
        onClose={() => setDeleteSchemaVisible(false)}
      />
    </>
  );
};

export default SchemaEdit;
