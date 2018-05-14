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
    this.forceUpdateHandler = this.forceUpdateHandler.bind(this);
  }

  componentWillMount() {
    window.img1 = "";
    window.img2 = "";
  }

  handleChange(event){
      this.setState({atomTxt: event.target.value});
  }

  handleModelChange (model) {
    this.setState({
      content: model
    });
  }

  forceUpdateHandler(){
    var req1 = "<p>- - - FromDevice p1</p><p>FromDevice - - Classifier 12/0800 -</p><p>Classifier 1 - ToDevice p2</p><p>Classifier 0 - Strip 14</p><p>Strip - - CheckIPHeader -</p><p>CheckIPHeader - - ReadState -</p><p>ReadState 1 - ToDevice p2</p><p>ReadState 0 - SessionCheck -</p><p>SessionCheck 1 - ToDevice p2</p><p>SessionCheck 2 - Discard -</p><p>SessionCheck 0 - WriteState -</p><p>WriteState - - ToDevice p2</p><p>- - - FromDevice p2</p><p>FromDevice - - ToDevice p1</p>"
    var req2 = "<p>- - - FromDevice p1</p><p>FromDevice - - Classifier 12/0800 -</p><p>Classifier 1 - ToDevice p2</p><p>Classifier 0 - Strip 14</p><p>Strip - - CheckIPHeader -</p><p>CheckIPHeader - - ReadState -</p><p>ReadState 1 - ToDevice p2</p><p>ReadState 0 - SessionCheck -</p><p>SessionCheck 1 - ToDevice p2</p><p>SessionCheck 2 - Log -</p><p>Log - - Discard -</p><p>SessionCheck 0 - WriteState -</p><p>WriteState - - ToDevice p2</p><p>- - - FromDevice p2</p><p>FromDevice - - ToDevice p1</p>"
    let cont = {
      atom_Txt : this.state.atomTxt,
      content : this.state.content,
    }
    if(req1 == cont.content){
      window.img1 = "http://192.168.0.13:4000/frontend/images/clickreq1.svg?" + new Date;
    }
    if(req2 == cont.content){
      window.img1 = "http://192.168.0.13:4000/frontend/images/clickreq2.svg?" + new Date;
    }
    if(cont.content == ""){
      window.img1 = ""
    }
    document.getElementsByTagName('change1').src = window.img1;
    console.log(document.getElementsByTagName('change1').src);
    this.forceUpdate();
  }


  handleClick(event){
    var req1 = "<p>- - - FromDevice p1</p><p>FromDevice - - Classifier 12/0800 -</p><p>Classifier 1 - ToDevice p2</p><p>Classifier 0 - Strip 14</p><p>Strip - - CheckIPHeader -</p><p>CheckIPHeader - - ReadState -</p><p>ReadState 1 - ToDevice p2</p><p>ReadState 0 - SessionCheck -</p><p>SessionCheck 1 - ToDevice p2</p><p>SessionCheck 2 - Discard -</p><p>SessionCheck 0 - WriteState -</p><p>WriteState - - ToDevice p2</p><p>- - - FromDevice p2</p><p>FromDevice - - ToDevice p1</p>"
    var req2 = "<p>- - - FromDevice p1</p><p>FromDevice - - Classifier 12/0800 -</p><p>Classifier 1 - ToDevice p2</p><p>Classifier 0 - Strip 14</p><p>Strip - - CheckIPHeader -</p><p>CheckIPHeader - - ReadState -</p><p>ReadState 1 - ToDevice p2</p><p>ReadState 0 - SessionCheck -</p><p>SessionCheck 1 - ToDevice p2</p><p>SessionCheck 2 - Log -</p><p>Log - - Discard -</p><p>SessionCheck 0 - WriteState -</p><p>WriteState - - ToDevice p2</p><p>- - - FromDevice p2</p><p>FromDevice - - ToDevice p1</p>"
    let cont = {
      atom_Txt : this.state.atomTxt,
      content : this.state.content,
    }
    var ele_list = ["FromDevice","Classifier","Strip","CheckIPHeader","EtherEncap","Discard","FullNoteQueue","ToDevice"]
    var atom_list = ["ReadState","SessionCheck","WriteState"]
    let udf = {
      Element_name : "Firewall",
      Atom_name : atom_list,
    }
    var udf_list = [udf]
    let param = {
      //Vnf_config : "FromDevice(eth0) -> cla :: Classifier(12/0800 -);cla[1] -> ToDevice(eth1);cla[0] -> Strip(14) -> CheckIPHeader -> ipcla :: IPClassifier(tcp -);ipcla[1] -> ToDevice(eth1);ipcla[0] -> fir :: Firewall;fir[1] -> Discard;fir[0] -> ToDevice(eth1);",
      Vnf_config : "icla :: Classifier(12/0800, -);ifi :: Firewall;to_extern :: FullNoteQueue -> ToDevice(p2);FromDevice(p1) -> icla;icla[1] -> to_extern;icla[0] -> Strip(14)-> CheckIPHeader-> ifi;ifi[1] -> Discard;ifi[0] -> EtherEncap(0x0800, fa:fe:ca:5d:97:6c, ee:43:35:3d:55:7c) ->to_extern;FromDevice(p2) -> FullNoteQueue -> ToDevice(p1);",
      Element_list : ele_list,
      User_defined_element : udf_list,
    }
    let data1 = {
      Id : 1,
      Jsonrpc : 2.0,
      Method : "create",
      Params : param,
    }
    let atom_list2 = ["Log"]
    let udf2 = {
      Element_name : "Log",
      Atom_name : atom_list2,
    }
    var udf_list2 = [udf2]
    var ele_list2 = ["FromDevice","Classifier","Strip","CheckIPHeader","EtherEncap","Discard","FullNoteQueue","ToDevice","Firewall"]
    let param2 = {
      //Vnf_config : "FromDevice(eth0) -> cla :: Classifier(12/0800 -);cla[1] -> ToDevice(eth1);cla[0] -> Strip(14) -> CheckIPHeader -> ipcla :: IPClassifier(tcp -);ipcla[1] -> ToDevice(eth1);ipcla[0] -> fir :: Firewall;fir[1] -> Log -> Discard;fir[0] -> ToDevice(eth1);",
      Vnf_config : "icla :: Classifier(12/0800, -);ifi :: Firewall;to_extern :: FullNoteQueue -> ToDevice(p2);FromDevice(p1) -> icla;icla[1] -> to_extern;icla[0] -> Strip(14)-> CheckIPHeader-> ifi;ifi[1] -> Log -> Discard;ifi[0] -> EtherEncap(0x0800, fa:fe:ca:5d:97:6c, ee:43:35:3d:55:7c) ->to_extern;FromDevice(p2) -> FullNoteQueue -> ToDevice(p1);",
      Element_list : ele_list2,
      User_defined_element : udf_list2,
    }
    let data2 = {
      Id : 1,
      Jsonrpc : 2.0,
      Method : "update",
      Params : param2,
    }
    console.log(cont.content)
    if(req1 == cont.content){
      fetch('http://192.168.0.13:4000/update',{
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
        window.img2 = "http://192.168.0.13:4000/frontend/images/clickreq1.svg?" + new Date;
        document.getElementsByTagName('change2').src = window.img2;
        console.log(document.getElementsByTagName('change2').src);
        this.forceUpdate();
      }
    if (req2 == cont.content) {
      fetch('http://192.168.0.13:4000/update',{
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
        window.img2 = "http://192.168.0.13:4000/frontend/images/clickreq2.svg?" + new Date;
        document.getElementsByTagName('change2').src = window.img2;
        console.log(document.getElementsByTagName('change2').src);
        this.forceUpdate();
      }
    if (cont.content == "") {
      fetch('http://192.168.0.13:4000/delete',{
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
        window.img2 = "";
        document.getElementsByTagName('change2').src = window.img2;
        console.log(document.getElementsByTagName('change2').src);
        this.forceUpdate();
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
          <div className="col-md-6 expertice-left-grid wow fadeInLeft" data-wow-delay="0.4s">
            <div className="expertise-head">
              <h3>loaded Atom graph</h3>
            </div>
            <div>
              <img id="change2" src={window.img2} alt="" />
            </div>
          </div>
        </div>
        <div className="expertice-grids">
          <div className="col-md-6 expertice-left-grid wow fadeInRight" data-wow-delay="0.4s">
            <div className="expertise-head">
              <h3>Atom graph Design</h3>
            </div>
            <div>
              <img id="change1" src={window.img1} alt=""/>
            </div>
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
              <a id="refresh" className="leran-more" onClick={this.forceUpdateHandler}>refresh</a>
              <a id="submit" className="leran-more" onClick={this.handleClick.bind(this)}>submit</a>
            </div>
          </div>
          <div className="clearfix"> </div> 
        </div>
    </div>

    );
  }
}

export default App;
