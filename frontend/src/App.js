import React from "react";
import "./App.css";
import { SnackbarProvider } from "notistack";
import ParamView from "./components/ParamView";

function App() {
  return (
    <SnackbarProvider maxSnack={1}>
      <div id="app" className="App">
        <ParamView />
      </div>
    </SnackbarProvider>
  );
}

export default App;
