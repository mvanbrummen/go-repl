import React, { Component } from 'react';
import brace from 'brace';
import 'brace/mode/golang';
import 'brace/mode/text';
import 'brace/theme/solarized_dark';
import 'brace/theme/terminal';
import AceEditor from 'react-ace';
import { FaPlay, FaTimes } from 'react-icons/fa';
import { getVersion, executeCode } from './util/Gateway'

export default class App extends Component {
  state = {
    editorValue: "",
    result: ""
  }

  executeCodeSubmit = (e) => {
    e.preventDefault();

    executeCode(this.state.editorValue).then((resp) => {
      console.log(resp.result)
      this.setState({
        result: resp.result
      });
    });
  }

  clearEditor = (e) => {
    this.setState({
      editorValue: ""
    });
  }

  componentDidMount() {
    getVersion().then((v) => {
      console.log(v);
      this.setState({
        result: [v.result]
      });
    });
  }

  onChange = (value, e) => {
    this.setState({
      editorValue: value
    });
  }

  render() {
    return (
      <div className="editor">
        <h1 className="brand">gorepl</h1>
        <button id="playBtn" onClick={this.executeCodeSubmit}><FaPlay /></button>
        <button id="clearBtn" onClick={this.clearEditor}><FaTimes /></button>

        <AceEditor
          theme="solarized_dark"
          mode="golang"
          name="replEditor"
          fontSize={16}
          editorProps={{ $blockScrolling: true }}
          value={this.state.editorValue}
          onChange={this.onChange}
          width="100%"
          height="400px"
        />
        <AceEditor
          theme="terminal"
          mode="text"
          name="resultEditor"
          editorProps={{ $blockScrolling: true }}
          value={this.state.result}
          width="100%"
          height="200px"
          readOnly={true}
          highlightActiveLine={false}
          fontSize={16}
        />
      </div>
    );
  }
}
