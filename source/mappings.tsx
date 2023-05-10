import React from 'react';
import Help from "./components/Help.js";
import WhoAmI from "./components/WhoAmI.js";

interface IMapping {
  instructions: string;
  component: (params: (string | undefined)[]) => React.ReactElement;
}

const mappings: Map<string, IMapping> = new Map();

mappings.set('help', {
  instructions: 'Type "help" to see this list again.',
  component: () => { return <Help mappings={mappings} />; }
});
mappings.set('whoami', {
  instructions: 'Type "whoami" to see who you are.',
  component: () => { return <WhoAmI />; }
});

export default mappings;