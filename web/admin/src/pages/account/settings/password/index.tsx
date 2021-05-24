import React, { useState, useEffect } from "react";
import { message } from 'antd';
import ProForm, { ProFormText } from "@ant-design/pro-form";
import { LockOutlined } from "@ant-design/icons";

export default (): React.ReactNode => {
  return (
    <ProForm
      onFinish={async () => {
        await waitTime(2000);
        message.success('提交成功');
      }}
      submitter={{
        searchConfig: {
          submitText: '修改密码',
        },
        render: (_, dom) => dom.pop(),
      }}
    >
      <ProFormText.Password
        name="oldPwd"
        label="旧密码"
        width="md"
        placeholder="请输入旧密码"
        fieldProps={{
          prefix: <LockOutlined />
        }}
        rules={[
          {
            required: true,
            message: '请输入旧密码!',
          },
        ]}
      />
      <ProFormText.Password
        name="newPwd"
        label="新密码"
        width="md"
        placeholder="请输入新密码"
        fieldProps={{
          prefix: <LockOutlined />
        }}
        rules={[
          {
            required: true,
            message: '请输入新密码!',
          },
        ]}
      />
      <ProFormText.Password
        name="newPwd2"
        label="确认密码"
        width="md"
        placeholder="请二次输入新密码"
        fieldProps={{
          prefix: <LockOutlined />
        }}
        rules={[
          {
            required: true,
            message: '请二次输入新密码!',
          },
        ]}
      />
    </ProForm>
  )
}
