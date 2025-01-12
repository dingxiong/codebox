import { useState } from 'react';
import { Button, Space, message, Select, Input, Alert } from 'antd';
import { JsonViewer } from '@textea/json-viewer';
import { beautifyJson, isJson } from '../utils';
import { CallES } from '../../../wailsjs/go/main/App';
import styles from './request.module.css';

const { TextArea } = Input;

enum HttpMethod {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
}

enum Environment {
  STAGING = 'staging',
  PRODUCTION = 'production',
}

const HttpRequestBoard = ({
  setResponse,
}: {
  setResponse: (a: string) => void;
}) => {
  const [environment, setEnvironment] = useState<Environment>(
    Environment.STAGING
  );
  const [method, setMethod] = useState<HttpMethod>(HttpMethod.GET);
  const [path, setPath] = useState<string>('/_cat/indices?v');
  const [payload, setPayload] = useState<string | undefined>();
  const [isPayloadJson, setIsPayloadJson] = useState<boolean>(true);

  const onEnvChange = (value: string) => {
    setEnvironment(value as Environment);
  };

  const onMethodChange = (value: string) => {
    setMethod(value as HttpMethod);
  };

  const onPathChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { value: inputValue } = e.target;
    setPath(inputValue);
  };

  const sendRequest = () => {
    // message.info(
    //   `Env: ${environment}. Method: ${method} Path: ${path}. Payload: ${payload}`
    // );
    CallES(environment, method, path).then(setResponse);
  };

  const onPayloadChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const v = e.target.value.trim();
    setPayload(v);
    setIsPayloadJson(v === undefined || v === '' || isJson(v));
  };

  const beautifyPayload = () => {
    if (payload) {
      setPayload(beautifyJson(payload));
    }
  };

  const selectBefore = (
    <Select defaultValue={method} onChange={onMethodChange}>
      {Object.keys(HttpMethod).map((v) => (
        <Select.Option value={v} key={v}>
          {v}
        </Select.Option>
      ))}
    </Select>
  );

  const EnvOptions = Object.values(Environment).map((v) => {
    return { value: v, label: v };
  });

  return (
    <>
      <div>
        <Space.Compact style={{ width: '100%' }}>
          <Select
            defaultValue={environment}
            style={{ width: 180 }}
            onChange={onEnvChange}
            options={EnvOptions}
          />
          <Input
            addonBefore={selectBefore}
            defaultValue={path}
            onChange={onPathChange}
          />
          <Button type="primary" onClick={sendRequest}>
            Send
          </Button>
        </Space.Compact>
        <TextArea
          showCount
          maxLength={10000}
          placeholder="Request body..."
          autoSize={{ minRows: 10, maxRows: 20 }}
          onChange={onPayloadChange}
          value={payload}
        />
        {isPayloadJson && payload && (
          <Space direction="horizontal" style={{ width: '100%' }}>
            <Button onClick={beautifyPayload}>Beautify</Button>
          </Space>
        )}
        {!isPayloadJson && (
          <Space direction="vertical" style={{ width: '100%' }}>
            <Alert
              type="error"
              message="Payload is not a valid json!"
              banner
            />
          </Space>
        )}
      </div>
    </>
  );
};

const HttResponseBoard = ({ payload }: { payload: string }) => {
  if (!payload) {
    return <></>;
  }

  return (
    <div className={styles.floatdiv}>
      {isJson(payload) ? (
        <JsonViewer indentWidth={2} value={JSON.parse(payload)} />
      ) : (
        <TextArea
          autoSize={{ minRows: 10, maxRows: 25 }}
          value={payload}
          wrap="off"
          readOnly
        ></TextArea>
      )}
    </div>
  );
};

export const RequestAndResponse = () => {
  const [response, setResponse] = useState<string>('');
  return (
    <div>
      <HttpRequestBoard setResponse={setResponse}></HttpRequestBoard>
      <HttResponseBoard payload={response} />
    </div>
  );
};
