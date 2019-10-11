import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import { fetchAdminLeavePending, deleteRequestLeave } from "../../../store/Actions/adminActions";
import HeaderAdmin from "../../../pages/menu/HeaderAdmin";
import Loading from "../../../components/Loading";
import Footer from "../../../components/Footer";
import "./style.css";
import {
  Layout, Table, Modal, Button, Input, Icon, Divider, Popconfirm, message
} from "antd";
const { Content } = Layout;
let data;

class AdminPendingPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      visible: false,
      user: null,
      data: this.props.leave,
      filterDropdownVisible: false,
      filterDropdownNameVisible: false,
      filtered: false,
      searchText: "",
      searchID: ""
    };
  }

  componentWillMount() {
    console.log(
      " ----------------- Admin-List-Pending-Request ----------------- "
    );
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.leave !== this.props.leave) {
      this.setState({ data: nextProps.leave });
    }
    data = nextProps.leave;
  }

  componentDidMount() {
    if (!localStorage.getItem("token")) {
      this.props.history.push("/");
    } else if (localStorage.getItem("role") !== "admin") {
      this.props.history.push("/");
    }
    this.props.fetchAdminLeavePending();
  }

  onSearch = () => {
    const { searchText } = this.state;
    const reg = new RegExp(searchText, "gi");
    this.setState({
      filterDropdownNameVisible: false,
      filtered: !!searchText,
      data: data
        .map(record => {
          const match = record.name.match(reg);
          if (!match) {
            return null;
          }
          return {
            ...record,
            name: (
              <span>
                {record.name
                  .split(
                    new RegExp(`(?<=${searchText})|(?=${searchText})`, "i")
                  )
                  .map(
                    (text, i) =>
                      text.toLowerCase() === searchText.toLowerCase() ? (
                        <span key={i}>{text}</span>
                      ) : (
                          text
                        ) // eslint-disable-line
                  )}
              </span>
            )
          };
        })
        .filter(record => !!record)
    });
  };

  onSearchID = () => {
    const { searchID } = this.state;
    const reg = new RegExp(searchID, "gi");
    this.setState({
      filterDropdownIDVisible: false,
      filtered: !!searchID,
      data: data
        .map(record => {
          const match = String(record.id).match(reg);
          if (!match) {
            return null;
          }
          return {
            ...record,
            ID: (
              <span>
                {this.state.data.map(
                  (text, i) =>
                    String(text.id) === searchID ? (
                      <span key={i} className="highlight">
                        {text}
                      </span>
                    ) : (
                        text
                      ) // eslint-disable-line
                )}
                }
              </span>
            )
          };
        })
        .filter(record => !!record)
    });
  };

  onInputChangeID = e => {
    this.setState({
      searchID: e.target.value
    });
  };

  onInputChangeName = e => {
    this.setState({
      searchText: e.target.value
    });
  };

  showDetail = record => {
    this.setState({
      visible: true,
      user: record
    });
  };

  onSelectChange = selectedRowKeys => {
    console.log("selected row: ", selectedRowKeys);
  };

  handleCancel = () => {
    this.setState({ visible: false });
  };

  onShowSizeChange(current, pageSize) {
    console.log(current, pageSize);
  }

  deleteRequestLeave = (leaves, id) => {
    this.props.deleteRequestLeave(leaves, id);
  };

  render() {
    const { visible, loading } = this.state;
    const columns = [
      {
        title: "Request ID",
        dataIndex: "id",
        key: "id",
        width: 100,
        filterDropdown: (
          <div className="custom-filter-dropdown-id">
            <Input
              type="number"
              ref={ele => (this.searchInput = ele)}
              placeholder="Search request id"
              value={this.state.searchID}
              onChange={this.onInputChangeID}
              onPressEnter={this.onSearchID}
            />
            <Button type="primary" onClick={this.onSearchID}>
              Search
            </Button>
          </div>
        ),
        filterIcon: (
          <Icon
            type="search"
            style={{ color: this.state.filtered ? "#108ee9" : "#aaa" }}
          />
        ),
        filterDropdownIDVisible: this.state.filterDropdownIDVisible,
        onFilterDropdownVisibleChange: visible => {
          this.setState(
            {
              filterDropdownIDVisible: visible
            },
            () => this.searchInput && this.searchInput.focus()
          );
        }
      },
      {
        title: "Employee ID",
        dataIndex: "employee_number",
        key: "employee_number",
        width: 100
      },
      {
        title: "Name",
        dataIndex: "name",
        key: "name",
        width: 200,
        filterDropdown: (
          <div className="custom-filter-dropdown-name">
            <Input
              ref={ele => (this.searchInput = ele)}
              placeholder="Search name"
              value={this.state.searchText}
              onChange={this.onInputChangeName}
              onPressEnter={this.onSearch}
            />
            <Button type="primary" onClick={this.onSearch}>
              Search
            </Button>
          </div>
        ),
        filterIcon: (
          <Icon
            type="search"
            style={{ color: this.state.filtered ? "#108ee9" : "#aaa" }}
          />
        ),
        filterDropdownNameVisible: this.state.filterDropdownNameVisible,
        onFilterDropdownVisibleChange: visible => {
          this.setState(
            {
              filterDropdownNameVisible: visible
            },
            () => this.searchInput && this.searchInput.focus()
          );
        }
      },

      {
        title: "Position",
        dataIndex: "position",
        key: "position",
        width: 150
      },
      {
        title: "Email",
        dataIndex: "email",
        key: "email",
        width: 150
      },
      {
        title: "Type Of Leave",
        dataIndex: "type_name",
        key: "type_name",
        width: 150
      },
      {
        title: "From",
        dataIndex: "date_from",
        key: "date_from",
        width: 120
      },
      {
        title: "To",
        dataIndex: "date_to",
        key: "date_to",
        width: 120
      },
      {
        title: "Action",
        key: "action",
        width: 200,
        render: (value, record) => (
          <span>
            <Button type="primary" onClick={() => this.showDetail(record)}>
              Detail
            </Button>
            <Divider type="vertical" />
            <Popconfirm
              placement="top"
              title={"Are you sure delete this leave request?"}
              onConfirm={() => {
                this.deleteRequestLeave(this.props.leave, record.id);
                message.success("Leave request has been delete!");
              }}
              okText="Yes"
              cancelText="No"
            >
              <Button type="danger">Delete</Button>
            </Popconfirm>
          </span>
        )
      }
    ];

    if (this.props.loading) {
      return <Loading />;
    } else {
      return (
        <Layout>
          <HeaderAdmin />
          <Content
            className="container"
            style={{
              display: "flex",
              margin: "20px 16px 0",
              justifyContent: "center",
              paddingBottom: "606px"
            }}
          >
            <div style={{ padding: "20px", backgroundColor: "#fff", width: "1700px" }}>
              <h3>List of Pending Request</h3>
              <Table
                columns={columns}
                dataSource={this.state.data}
                rowKey={record => record.id}
                onRowClick={this.onSelectChange}
                pagination={{
                  className: "my-pagination",
                  defaultCurrent: 1,
                  defaultPageSize: 5,
                  total: `${this.state.data && this.state.data.length}`,
                  // showSizeChanger: this.onShowSizeChange
                }}
              />
            </div>

            <Modal
              visible={visible}
              title="Detail Leave Request Pending"
              onCancel={this.handleCancel}
              style={{ top: "20" }}
              bodyStyle={{ padding: "0" }}
              width="600px"
              footer={[
                <Button
                  key="cancel"
                  loading={loading}
                  onClick={this.handleCancel}
                >
                  Return
                </Button>
              ]}
            >
              <div style={{ padding: 10, background: "#fff" }}>
                <table>
                  <thead></thead>
                  <tbody>
                    <tr>
                      <td>ID</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.id}</td>
                    </tr>
                    <tr>
                      <td>Name</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.name}</td>
                    </tr>
                    <tr>
                      <td>Gender</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.gender}</td>
                    </tr>
                    <tr>
                      <td>Email</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.email}</td>
                    </tr>
                    <tr>
                      <td>Type Of Leave</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.type_name}</td>
                    </tr>
                    {this.state.user && this.state.user.reason !== "" ?
                      (
                        <tr>
                          <td>Reason</td>
                          <td>&nbsp;:</td>
                          <td>&nbsp;{this.state.user && this.state.user.reason}</td>
                        </tr>
                      ) :
                      (
                        <tr>
                          <td>Reason</td>
                          <td>&nbsp;:</td>
                          <td>&nbsp;-</td>
                        </tr>
                      )
                    }
                    <tr>
                      <td>From</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.date_from}</td>
                    </tr>
                    <tr>
                      <td>To</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.date_to}</td>
                    </tr>
                    {this.state.user && this.state.user.half_dates !== "{}" ?
                      (
                        <tr>
                          <td>Half Day</td>
                          <td>&nbsp;:</td>
                          <td>&nbsp;{this.state.user && this.state.user.half_dates.substring(1, this.state.user.half_dates.length - 1)}</td>
                        </tr>
                      ) :
                      (
                        <tr>
                          <td>Half Day</td>
                          <td>&nbsp;:</td>
                          <td>&nbsp;-</td>
                        </tr>
                      )
                    }
                    <tr>
                      <td>Back On</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.back_on}</td>
                    </tr>
                    <tr>
                      <td>Total Leave</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.total} day(s)</td>
                    </tr>
                    <tr>
                      <td>Leave Balance</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.leave_remaining} days</td>
                    </tr>
                    <tr>
                      <td>Contact Address</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.contact_address}</td>
                    </tr>
                    <tr>
                      <td>Contact Number</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.contact_number}</td>
                    </tr>
                    <tr>
                      <td>Status</td>
                      <td>&nbsp;:</td>
                      <td>&nbsp;{this.state.user && this.state.user.status}</td>
                    </tr>
                    {this.state.user && this.state.user.notes !== "" ?
                      (
                        <tr>
                          <td>Notes</td>
                          <td>&nbsp;:</td>
                          <td>&nbsp;{this.state.user && this.state.user.notes}</td>
                        </tr>
                      ) :
                      (
                        <tr>
                          <td>Notes</td>
                          <td>&nbsp;:</td>
                          <td>&nbsp;-</td>
                        </tr>
                      )
                    }
                  </tbody>
                </table>
              </div>
            </Modal>
          </Content>
          <Footer />
        </Layout>
      );
    }
  }
}

const mapStateToProps = state => ({
  loading: state.adminReducer.loading,
  leave: state.adminReducer.leaves
});

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      fetchAdminLeavePending,
      deleteRequestLeave
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AdminPendingPage);
