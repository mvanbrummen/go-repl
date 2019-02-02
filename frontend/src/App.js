import React, { Component } from 'react';
import brace from 'brace';
import 'brace/mode/golang';
import 'brace/theme/solarized_dark';
import { split as SplitEditor } from 'react-ace';
import { FaPlay } from 'react-icons/fa';

export default class App extends Component {
  render() {
    return (
      <div className="editor">
        <h1>gorepl</h1>
        <button id="playBtn"><FaPlay /></button>

        <SplitEditor
          theme="solarized_dark"
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
