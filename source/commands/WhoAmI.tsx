import React, { useEffect, useState } from 'react';
import { Box } from 'ink';
import Spinner from '../components/Spinner.js';
import KVP from '../components/KVP.js';
import UhOh from '../components/UhOh.js';
import api from '../lib/api.js';

interface IDebug {
  environment: string,
  scope: string,
  tenantId: string,
  tenantName: string
}

function delay(ms: number) {
  if (ms === 0) {
    return Promise.resolve();
  }
  return new Promise(resolve => setTimeout(resolve, ms));
}

export default () => {
  const [resp, setResp] = useState<IDebug | undefined>();
  const [error, setError] = useState<string | undefined>();

  useEffect(() => {
    delay(2000)
    .then(() => api('/debug', 'POST'))
    .then(
      ({ json }) => setResp(json),
      (err: Error) => setError(err.message)
    );
  }, []);

  if (error) {
    return <UhOh text={error} />;
  } else if (resp) {
    return <Box flexDirection="column">
      <KVP width={20} label="Workspace Name" value={resp.tenantName} />
      <KVP width={20} label="Workspace ID" value={resp.tenantId} />
      <KVP width={20} label="API Key Environment" value={resp.environment} />
      <KVP width={20} label="API Key Scope" value={resp.scope} />
    </Box>
  } else {
    return <Spinner text="pondering your existence..." />
  }
}
