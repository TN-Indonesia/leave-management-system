import React from "react";
import "./style.css";
import { Spin } from "antd";

export default class LoadingSpinner extends React.Component {
  render() {
    return (
      <div
        style={{
          padding: 250,
          justifyContent: "center"
        }}
      >
        <Spin size="large" />
      </div>
    );
  }
}
