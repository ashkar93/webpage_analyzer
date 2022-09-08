import Axios from 'axios';
import './App.css';
import AppBarMain from './AppBar';
import WebPageAnalyzer from './components/WebPageAnalyzer';


Axios.defaults.baseURL = "http://localhost:8000/";

function App() {
  return (
    <div className="App">
      <AppBarMain/>
      <WebPageAnalyzer/>
    </div>
  );
}

export default App;
