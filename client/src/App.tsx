import React from 'react';
import './App.css';
import { InputBox } from './components/InputBox';


const App: React.FC = () => {
  return (
    <div className="App">
      
        <div className="header">
          Hello
        </div>
        
        <div className="content">
          <InputBox/>
        </div>

        <div className="footer">
          <footer>Test</footer>
        </div>
      
    </div>
  );
}

export default App;
