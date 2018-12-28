import React, { Component } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { adminEditLeaveBalances, handleEditBalance, saveEditBalanceUser } from "../../../store/Actions/adminActions";
import HeaderAdmin from "../../../pages/menu/HeaderAdmin";
import Loading from "../../../components/Loading";
import Footer from "../../../components/Footer";
import "./style.css";
import { Layout, Button, Form, Input } from "antd";
const { Content } = Layout;
const FormItem = Form.Item;

class AdminEditBalancePage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      date: ""
    };

    this.saveEditBalanceUser = this.saveEditBalanceUser.bind(this);
    this.handleOnChange = this.handleOnChange.bind(this);
  }

  componentWillMount() {
    console.log(
      " ----------------- Admin-Edit-Balance ----------------- "
    );
  }

  componentDidMount() {
    if (localStorage.getItem("role") !== "admin") {
      this.props.history.push("/");
    }
    let employee_id = Number(
      this.props.history.location.pathname.split("/").pop()
    );

    this.props.adminEditLeaveBalances(employee_id);
  }

  saveEditBalanceUser = () => {
    console.log(this.props.balance)
    let employee_id = Number(
      this.props.history.location.pathname.split("/").pop()
    );

    this.props.saveEditBalanceUser(this.props.balance, employee_id, url => {
      this.props.history.push(url);
    });
  };

  handleOnChange = e => {
    let editEl = {
      ...this.props.balance[e.target.id],
      leave_remaining: Number(e.target.value)
    };
    let edit = [
      ...this.props.balance
    ]
    edit[e.target.id] = editEl
    console.log(edit)
    this.props.handleEditBalance(edit);
  };

  handleBlur() {
    console.log("blur");
  }

  handleFocus() {
    console.log("focus");
  }

  render() {
    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 }
      }
    };

    let elements = [];    
    {
      this.props.balance && this.props.balance.map((el, id) => {
        elements.push(
          <FormItem {...formItemLayout} label={el.type_name}>
            <Input
              type="number"
              id={id}
              name={el.type_name}
              defaultValue={el.leave_remaining}
              onChange={this.handleOnChange}
            />
          </FormItem>
        );
        return el
      })
    };

    if (this.props.loading) {
      return <Loading />;
    } else {
      return (
        <div>
          <Layout>
            <HeaderAdmin />
            <Content
              className="container"
              style={{
                display: "flex",
                margin: "24px 16px 0",
                justifyContent: "space-around",
                paddingBottom: "160px"
              }}
            >
              <div
                style={{
                  padding: 100,
                  paddingBottom: 50,
                  paddingTop: 50,
                  background: "#fff",
                  minHeight: 360
                }}
              >
                <h1> EDIT BALANCE </h1>
                <div>
                  <Form onSubmit={this.handleSubmit} className="login-form">
                    {elements}
                    <FormItem>
                      <Button
                        onClick={() => {
                          this.saveEditBalanceUser();
                        }}
                        htmlType="submit"
                        type="primary"
                      >
                        UPDATE
                      </Button>
                    </FormItem>
                  </Form>
                </div>
              </div>
            </Content>

            <Footer />
          </Layout>
        </div>
      );
    }
  }
}

const mapStateToProps = state => ({
  loading: state.adminReducer.loading,
  balance: state.adminReducer.balances
});

const WrappedEditUserForm = Form.create()(AdminEditBalancePage);

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      adminEditLeaveBalances,
      handleEditBalance,
      saveEditBalanceUser
    },
    dispatch
  );
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(WrappedEditUserForm);
