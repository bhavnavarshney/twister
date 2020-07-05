import React from 'react';
import './App.css';
import { SnackbarProvider } from 'notistack';
import HelloWorld from './components/HelloWorld';

function App() {
  return (
    <SnackbarProvider maxSnack={1}>
    <div id="app" className="App">
       <HelloWorld />
    </div>
    </SnackbarProvider>
  );
}

export default App;
