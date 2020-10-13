import React from 'react';
import { Switch, Route } from 'react-router-dom';

import Home from '../Home/Home';
import CreateCollage from '../CreateCollage/CreateCollage';
import DisplayCollage from '../DisplayCollage/DisplayCollage';
import PreviousCollage from '../PreviousCollage/PreviousCollage';

const Main = () => {
  return (
    <Switch> {/* The Switch decides which component to show based on the current URL.*/}
      <Route exact path='/' component={Home}></Route>
      <Route exact path='/create' component={CreateCollage}></Route>
      <Route exact path='/previous' component={PreviousCollage}></Route>
      <Route exact path='/display' component={DisplayCollage}></Route>
    </Switch>
  );
}

export default Main;