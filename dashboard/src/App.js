import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import './style.css';
import './animate.css';
import './bootstrap.css';
import 'froala-editor/js/froala_editor.pkgd.min.js';
import 'froala-editor/css/froala_style.min.css';
import 'froala-editor/css/froala_editor.pkgd.min.css';
import 'font-awesome/css/font-awesome.css';
import FroalaEditor from 'react-froala-wysiwyg';

class App extends Component {

  constructor(props) {
      super(props);
      this.state = {atomTxt: '',
        content: ''
      }
    this.handleModelChange = this.handleModelChange.bind(this);
  }

  handleChange(event){
      this.setState({atomTxt: event.target.value});
  }

  handleModelChange (model) {
    this.setState({
      content: model
    });
  }

  handleClick(event){
    var req1 = "<p>FromDevice - - Classifier 12/0800 -</p><p>Classifier 1 - ToDevice eth1</p><p>Classifier 0 - Strip 14</p><p>Strip - - CheckIPHeader -</p><p>CheckIPHeader - - IPClassifier tcp -</p><p>IPClassifier 1 - ToDevice eth1</p><p>IPClassifier 0 - ReadState -</p><p>ReadState 1 - ToDevice eth1</p><p>ReadState 0 - SessionCheck -</p><p>SessionCheck 1 - ToDevice eth1</p><p>SessionCheck 2 - Discard -</p><p>SessionCheck 0 - WriteState -</p><p>WriteState - - ToDevice eth1</p>"
    var req2 = "<p>FromDevice - - Classifier 12/0800 -</p><p>Classifier 1 - ToDevice eth1</p><p>Classifier 0 - Strip 14</p><p>Strip - - CheckIPHeader -</p><p>CheckIPHeader - - IPClassifier tcp -</p><p>IPClassifier 1 - ToDevice eth1</p><p>IPClassifier 0 - ReadState -</p><p>ReadState 1 - ToDevice eth1</p><p>ReadState 0 - SessionCheck -</p><p>SessionCheck 1 - ToDevice eth1</p><p>SessionCheck 2 - Log -</p><p>Log - - Discard -</p><p>SessionCheck 0 - WriteState -</p><p>WriteState - - ToDevice eth1</p>"
    let cont = {
      atom_Txt : this.state.atomTxt,
      content : this.state.content,
    }
    var ele_list = ["FromDevice","Classifier","ToDevice","Strip","CheckIPHeader","IPClassifier"]
    var atom_list = ["ReadState","SessionCheck","WriteState"]
    let udf = {
      Element_name : "Firewall",
      Atom_action : atom_list,
    }
    var udf_list = [udf]
    let param = {
      Vnf_config : "FromDevice(eth0) -> cla :: Classifier(12/0800 -);cla[1] -> ToDevice(eth1);cla[0] -> Strip(14) -> CheckIPHeader -> ipcla :: IPClassifier(tcp -);ipcla[1] -> ToDevice(eth1);ipcla[0] -> fir :: Firewall;fir[1] -> Discard;fir[0] -> ToDevice(eth1);",
      Element_list : ele_list,
      User_defined_element : udf_list,
    }
    let data1 = {
      Id : 1,
      Jsonrpc : 2.0,
      Method : "update",
      Params : param,
    }
    let udf2 = {
      Element_name : "Log",
      Atom_action : "Log",
    }
    var udf_list2 = [udf2]
    var ele_list2 = ["FromDevice","Classifier","ToDevice","Strip","CheckIPHeader","IPClassifier","Firewall"]
    let param2 = {
      Vnf_config : "FromDevice(eth0) -> cla :: Classifier(12/0800 -);cla[1] -> ToDevice(eth1);cla[0] -> Strip(14) -> CheckIPHeader -> ipcla :: IPClassifier(tcp -);ipcla[1] -> ToDevice(eth1);ipcla[0] -> fir :: Firewall;fir[1] -> Log -> Discard;fir[0] -> ToDevice(eth1);",
      Element_list : ele_list2,
      User_defined_element : udf_list2,
    }
    let data2 = {
      Id : 1,
      Jsonrpc : 2.0,
      Method : "update",
      Params : param2,
    }
    if(req1 == cont.content){
      fetch('http://localhost:4000/update',{
            method: 'POST',
            //mode: 'cors',
            headers: {
                //'Accept': 'application/json',
                'Content-Type' : 'text/plain'
            },
            body: JSON.stringify(data1)
        })
            .then((response) => response.text())
            .then((responseText) => {
                console.log(JSON.parse(responseText));
            })
        .catch(function (error) {
            console.log('request failed', error)
        })
      }
    if (req2 == cont.content) {
      fetch('http://localhost:4000/update',{
            method: 'POST',
            //mode: 'cors',
            headers: {
                //'Accept': 'application/json',
                'Content-Type' : 'text/plain'
            },
            body: JSON.stringify(data2)
        })
            .then((response) => response.text())
            .then((responseText) => {
                console.log(JSON.parse(responseText));
            })
        .catch(function (error) {
            console.log('request failed', error)
        })
      }
    if (cont.content == "") {
      fetch('http://localhost:4000/delete',{
            mode: 'cors',
            method: 'GET',
        })
            .then((response) => response.text())
            .then((responseText) => {
                console.log(JSON.parse(responseText));
            })
        .catch(function (error) {
            console.log('request failed', error)
        })
      }
  }


  render() {
    return (
      <div>
        <div className="bg">
          <div id="home" className="header wow bounceInDown" data-wow-delay="0.4s">
            <div className="top-header">
              <div className="contianer">
                <div className="logo">
                  <a href="#"><img src={require("./logo.png")}/></a>
                </div>
                <nav className="top-nav">
                  <ul className="top-nav">
                    <li><a href="#graph" className="scroll">Current Atom Graph</a></li>
                    <li><a href="#atomdesign" className="scroll">Atom Action Design</a></li>
                  </ul>
                </nav>
                <div className="clearfix"> </div>
              </div>
            </div>
          </div>
          <div className="banner wow fadeIn" data-wow-delay="0.5s">
            <div className="container">
              <div className="banner-info text-center">
                <h1>CLICK-UP PLATFORM</h1><br />
                <span> </span>
                <p>Towards Software Upgrades of Click-driven Stateful Network Elements</p>
              </div>
            </div>
          </div>
        </div>
      <div id="port" className="expertise">
         <div className="expertice-grids">
          <div className="col-md-4 expertice-left-grid wow fadeInLeft" data-wow-delay="0.4s">
            <div className="expertise-head">
              <h3>loaded Atom graph</h3>
            </div>
            <div>
              <img src={require("./click.png")}/>
            </div>
          </div>
        </div>
        <div className="expertice-grids">
          <div className="col-md-4 expertice-left-grid wow fadeInLeft" data-wow-delay="0.4s">
            <div className="expertise-head">
              <h3>Atom graph Design</h3>
            </div>
            <div>
              <img src={require("./click.png")}/>
              <br />
              <a id="refresh" className="leran-more">refresh</a>
            </div>
          </div>
        </div>
        <div className="expertice-grids">
          <div className="col-md-4 expertice-left-grid wow fadeInRight" data-wow-delay="0.4s">
            <div className="expertise-head">
              <h3>Atom Action Design</h3>
            </div>
            <div style={{width:'90%', height:'500px'}}>
              <FroalaEditor id="txt" style={{width:'90%', height:'95%'}} model={this.state.content} onModelChange={this.handleModelChange}></FroalaEditor>
              <br />
              <a id="submit" className="leran-more" onClick={this.handleClick.bind(this)}>submit</a>
            </div>
          </div>
          <div className="clearfix"> </div> 
        </div>
      </div>
    </div>

    );
  }
}

export default App;
