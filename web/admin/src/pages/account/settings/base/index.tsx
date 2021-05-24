import React, { useState, useEffect } from "react";
import { message } from 'antd';
import ProForm, { ProFormSelect, ProFormUploadButton } from "@ant-design/pro-form";

export default (): React.ReactNode => {
  return (
    <ProForm
      onFinish={async () => {
        await waitTime(2000);
        message.success('提交成功');
      }}
      submitter={{
        searchConfig: {
          submitText: '更新基本信息',
        },
        render: (_, dom) => dom.pop(),
      }}
    >
      <ProFormUploadButton
        name="avatar"
        label="头像"
        max={1}
        fieldProps={{
          name: 'file',
          listType: 'picture-card',
        }}
        action="/upload.do"
      />
      <ProFormSelect
        width="md"
        name="language"
        label="语言"
        rules={[
          {
            required: true,
            message: '请选择语言!',
          },
        ]}
        options={[
          { label: '中文', value: 'zh-CN' },
          { label: 'EN', value: 'en' }
        ]}
      />
    </ProForm>
  )
}
