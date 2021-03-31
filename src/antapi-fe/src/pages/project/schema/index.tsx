import { useParams, connect } from 'umi';
import type { Dispatch } from 'umi';
import React, { useEffect } from 'react';
import { Button, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import ProCard from '@ant-design/pro-card';
import { PageContainer } from '@ant-design/pro-layout';
import SchemaMenuList from './SchemaMenuList';
import './index.less';

interface SchemasProps {
  dispatch: Dispatch;
  schemas: Partial<API.Schema[]>;
  loading: boolean;
}

const SchemaList: React.FC<SchemasProps> = (props) => {
  const { projectId } = useParams<{projectId: string}>();
  const { dispatch, schemas, loading } = props;
  const isLoading = loading || !schemas;

  console.log(isLoading);

  useEffect(() => {
    dispatch({
      type: 'schema/getSchemas',
      payload: {
        projectId
      }
    });
  }, [projectId, dispatch]);

  return (
    <PageContainer
      className="schema-page-container"
      fixedHeader
      affixProps={{ offsetTop: 48 }}
      extra={
        <Space>
          <Button
            type="primary"
            onClick={() => {
              ctx.mr.createSchema();
            }}
          >
            <PlusOutlined />
            新建模型
          </Button>
        </Space>
      }
    >
      <ProCard split="vertical" gutter={[16, 16]} style={{ background: 'inherit' }}>
      {/* 模型菜单 */}
        <ProCard colSpan="220px" className="card-left" style={{ marginBottom: 0 }}>
          <SchemaMenuList />
        </ProCard>
      </ProCard>
    </PageContainer>
  );
};

export default connect(({ schema, loading }: ConnectState) => ({
  schema,
  loading: loading.models.schema,
}))(SchemaList);
