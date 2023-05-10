import React from 'react';
import Help from "./commands/Help.js";
import WhoAmI from "./commands/WhoAmI.js";

interface IMapping {
  instructions?: string;
  component: (params: (string | undefined)[]) => React.ReactElement;
}

const mappings: Map<string, IMapping> = new Map();

mappings.set('help', {
  component: () => { return <Help mappings={mappings} />; }
});
mappings.set('whoami', {
  instructions: 'Type "whoami" to see who you are.',
  component: () => { return <WhoAmI />; }
});

export default mappings;