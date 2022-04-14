import React, { Component } from "react";
import { Route, Switch } from "react-router-dom";

import Home from "./pages/Home";
import ArchiveViewer from "./pages/ArchiveViewer";
import ClipCollections from "./pages/ClipCollections";


// const electron = window.require('electron');
// const { ipcRenderer } = electron;

class App extends Component {
    render() {
        return (
            <>
                <Switch>
                    <Route path="/" exact={true} component={Home} />
                    <Route path="/archive-viewer" component={ArchiveViewer} />
                    <Route path="/clip-collections" component={ClipCollections} />
                </Switch>
            </>
        )
    }
};

export default App;