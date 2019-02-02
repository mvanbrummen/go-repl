import React, { Component } from 'react';
import brace from 'brace';
import 'brace/mode/golang';
import 'brace/theme/github';
import { split as SplitEditor } from 'react-ace';

export default class App extends Component {
  render() {
    return (
      <div className="editor">
        <SplitEditor
          theme="github"
          mode="golang"
          name="replEditor"
          editorProps={{ $blockScrolling: true }}
          splits={2}
          orientation="besides"
          width="100%"
          height="600px"
        />
      </div>
    );
  }
}
