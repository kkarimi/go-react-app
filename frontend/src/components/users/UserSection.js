import React from 'react';
import UserForm from './UserForm.js';
import UserList from './UserList.js';

class UserSection extends React.Component{
  render(){
    return (
      <div className='support panel panel-primary'>
        <div className='panel-heading'>
          <strong>Users</strong>
        </div>
        <div className='panel-body users'>
          <UserList {...this.props} />
          <UserForm {...this.props} />
        </div>
      </div>
    )
  }
}

UserSection.propTypes = {
  users: React.PropTypes.array.isRequired,
  setUserName: React.PropTypes.func.isRequired
}

export default UserSection