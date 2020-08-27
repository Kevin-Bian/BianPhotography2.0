import React from 'react';
import { Route, IndexRoute } from 'react-router';

/**
 * Import all page components here
 */
// import App from '../components/App/App';
import Home from '../components/Home/Home';
import CreateCollage from '../components/CreateCollage/CreateCollage';

/**
 * All routes go here.
 * Don't forget to import the components above after adding new route.
 */
export default (
    <div>
        <Route path="/" component={Home} />
        <Route path="/home" component={Home} />
        <Route path="/create" component={CreateCollage} />
    </div>
);