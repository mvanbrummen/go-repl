import React, { Component } from 'react';
import brace from 'brace';
import 'brace/mode/golang';
import 'brace/mode/ruby';
import 'brace/mode/java';
import 'brace/mode/javascript';
import 'brace/mode/python';
import 'brace/mode/text';
import 'brace/theme/solarized_dark';
import 'brace/theme/terminal';
import AceEditor from 'react-ace';
import { FaPlay, FaTimes } from 'react-icons/fa';
import { getVersion, executeCode } from './util/Gateway'

const langauges = [
  "golang",
  "ruby",
  "javascript",
  "python",
  "java",
];

const javaDefault = `
public class Main {
  public static void main(String[] args) {
      System.out.println("Hello, World!");
  }
}
`;

export default class App extends Component {
  state = {
    editorValue: "",
    result: "",
    language: langauges[0]
  }

  executeCodeSubmit = (e) => {
    e.preventDefault();

    this.setState({
      result: "Executing program..."
    });

    executeCode(this.state.language, this.state.editorValue).then((resp) => {
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
    getVersion(this.state.language).then((v) => {
      this.setState({
        result: v.result
      });
    });
  }

  onChange = (value, e) => {
    this.setState({
      editorValue: value
    });
  }

  onLanguageChange = (e) => {
    let language = e.target.value;

    let editorValue = this.state.editorValue;
    if (language === "java") {
        editorValue = javaDefault;
    }
    getVersion(language).then((v) => {
      this.setState({
        editorValue: editorValue,
        result: v.result,
        language: language
      });
    });
  }

  render() {
    return (
      <div className="editor">
        <div className="container">
          <h1 className="brand">gorepl</h1>
          <button className="site" id="playBtn" onClick={this.executeCodeSubmit}><FaPlay /></button>
          <button id="clearBtn" className="site" onClick={this.clearEditor}><FaTimes /></button>
          <select onChange={this.onLanguageChange}>
            {
              langauges.map((language, i) =>
                <option key={i} value={language}>{language}</option>
              )
            }
          </select>
        </div>

        <AceEditor
          theme="solarized_dark"
          mode={this.state.language}
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
          height="250px"
          readOnly={true}
          highlightActiveLine={false}
          fontSize={16}
        />
      </div>
    );
  }
}
