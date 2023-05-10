import React from 'react';
import Help from "./Help.js";

interface IMapping {
  instructions: string;
  component: any;
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
  if (mapping) {
    return mapping.component();
  } else {
    return <Help mappings={mappings} />;
  }
}
