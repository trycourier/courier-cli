#!/usr/bin/env node
import React from 'react';
import { render} from 'ink';
import Help from "./Help.js";
import Counter from "./Counter.js";
import Error from "./Error.js";

const args = (argv: string[]): string[] => {
  if (!argv || !argv.length || !argv[0]) {
    return [];
  } else if (argv[0].endsWith('node')) {
    const [,, ...args] = argv;
    return args;
  } else {
    const [, ...args] = argv;
    return args;
  }
}

const CLI = () => {
  const params = args(process.argv)
  if (!params.length){
    return <Help />;
  }

  switch(params[0]) {
    case 'counter':
      return <Counter />;
    case 'error':
      return <Error />;
    default:
      return <Help />;
  }
}

render(<CLI />);
