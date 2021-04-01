import { useParams, connect } from 'umi';
import type { Dispatch } from 'umi';
import React, { useEffect, useState } from 'react';
import { Layout, Button, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import ProCard from '@ant-design/pro-card';
import { PageContainer } from '@ant-design/pro-layout';
import SchemaMenuList from './SchemaMenuList';
import SchemaContent from './SchemaContent';
import './index.less';

interface SchemasProps {
  dispatch: Dispatch;
  schemas: Partial<API.Schema[]>;
  loading: boolean;
}

const SchemaList: React.FC<SchemasProps> = (props) => {
  const { projectId } = useParams<{projectId: string}>();
  const { dispatch, schemas } = props;
  const [ currentSchemaId, setCurrentSchemaId ] = useState<string>('');
  // const isLoading = loading || !schemas;
  // const menuRef = useRef<API.Schema[]>(schemas || []);

  useEffect(() => {
    dispatch({
      type: 'schema/getSchemas',
      payload: {
        projectId
      }
    });
  }, [projectId, dispatch]);

  const onSelectSchema = (schemaId: string) => {
    // menuRef.current.onSelectSchema(schemaId);
    setCurrentSchemaId(schemaId);
  }

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
              console.log("todo create")
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
          <SchemaMenuList
            // ref={menuRef}
            currentSchemaId={currentSchemaId}
            schemas={schemas}
            onSelect={onSelectSchema}
          />
        </ProCard>
        {/* 模型字段 */}
        <Layout className="schema-layout">
          <SchemaContent currentSchema={currentSchemaId ? schemas?.find( element => element.id === currentSchemaId) : schemas[0]} />
        </Layout>
      </ProCard>
    </PageContainer>
  );
};

export default connect(
  ({
    schema,
    loading
  }: {
    schema: PageState;
    loading: { effects: Record<string, boolean> };
  }) => ({
    schemas: schema.schemas,
    loading: loading.effects['schema/getSchemas'],
  }),
)(SchemaList);
