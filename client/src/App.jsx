import React from 'react';
import './App.css';
import { Header } from './components/Header';
import { InputBox } from './components/InpuxBox';

function App() {
  return (
    <div className="App">
      <div className="col1"></div>
      <div className="col2">
        <div className="header">
            <Header/>
          </div>
          
          <div className="content">
            <InputBox/>
          </div>

          <div className="footer">
            <footer>Test</footer>
          </div>
      </div>
      <div className="col3"></div> 
    </div>
  );
}

export default App;
