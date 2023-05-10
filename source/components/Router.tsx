import React from 'react';
import Help from "../commands/Help.js";

interface IMapping {
  instructions?: string;
  component: (params: (string | undefined)[]) => React.ReactElement;
}

type Props = {
  args: string[];
  mappings: Map<string, IMapping>;
}

export default ({args, mappings}: Props) => {
  if (!args.length || !args[0]) {
    return <Help mappings={mappings} />;
  }

  const mapping = mappings.get(args[0]); 
  const params = [, ...args];

  if (mapping) {
    return mapping.component(params);
  } else {
    return <Help mappings={mappings} />;
  }
}
