import React, { useState, useCallback } from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import Gallery from "react-photo-gallery";
import { photos } from "../../Home/photos";
import Carousel, { Modal, ModalGateway } from "react-images";
import Button from '@material-ui/core/Button';

class SubmitForm extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      collageID: '',
      imageURL: '',
      imageName: '',
      imageDescription: '',
      collageSearchID: ''
    }

  }

  handleCollageIDChange = (e) => {
    this.setState({ collageID: e.target.value });
  };

  handleImageURLChange = (e) => {
    this.setState({ imageURL: e.target.value });
  };

  handleImageNameChange = (e) => {
    this.setState({ imageName: e.target.value });
  };

  handleImageDescriptionChange = (e) => {
    this.setState({ imageDescription: e.target.value });
  };

  handleCollageSearchChange = (e) => {
    this.setState({ collageSearchID: e.target.value });
  };

  onSubmit = (e) => {
    let payload = {
      CollageID: this.state.collageID,
      Name: this.state.imageName,
      Link: this.state.imageURL,
      Description: this.state.imageDescription
    }
    fetch('http://localhost:8080/apiv2/newphoto', {
    method: 'POST',
    body: JSON.stringify(payload)
  })
    return false
  }
  render() {
    return (

      <div className="card bg-light">
        <article className="card-body mx-auto" style={{ width: "500px" }}>
          <form>
          <h1 class="display-5">Add Photo To Collage</h1>
            <div className="form-group">
              <label for="collageID">Collage ID</label>
              <input type="text" className="form-control" id="collageID" placeholder="Enter an ID" value={this.state.collageID} onChange={this.handleCollageIDChange}></input>
            </div>
            <div className="form-group">
              <label for="imageURL">Image URL</label>
              <input type="text" className="form-control" id="imageURL" placeholder="Provide the image URL" value={this.state.imageURL} onChange={this.handleImageURLChange}></input>
            </div>
            <div className="form-group">
              <label for="imageName">Image Name</label>
              <input type="text" className="form-control" id="imageName" placeholder="Give your image a name" value={this.state.imageName} onChange={this.handleImageNameChange}></input>
            </div>
            <div className="form-group">
              <label for="imageDescription">Image Description</label>
              <input type="text" className="form-control" id="imageDescription" placeholder="Give a brief description" value={this.state.imageDescription} onChange={this.handleImageDescriptionChange}></input>
            </div>
            <button type="button" className="btn btn-primary" onClick={this.onSubmit}>Submit</button>
          </form>
        </article>

        <div className="card bg-light">
        <article className="card-body mx-auto" style={{ width: "500px" }}>
          <form>
            <h1 class="display-5">View Collage</h1>
            <div className="form-group">
              <label for="collageID">Enter the Collage ID</label>
              <input type="text" className="form-control" id="collageID" placeholder="Enter an ID" value={this.state.collageSearchID} onChange={this.handleCollageSearchChange}></input>
            </div>
            <Button variant="contained" color="secondary" href={"./display?id=" + this.state.collageSearchID}>View Collage </Button>
          </form>
        </article>
        </div>
        <Gallery photos={photos} margin={2} targetRowHeight={500} />

      </div>
    );
  }

}

export default SubmitForm