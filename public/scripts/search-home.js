var SearchPanel = React.createClass({displayName: "SearchPanel",
  switchToCreate: function() {
    console.log("create from search")
  },
  render: function() {
    return (
      React.createElement("div", {className: "search-panel text-center"}, 
        React.createElement("div", {className: "container"}, 
          React.createElement(SearchPage, {onAddService: this.switchToCreate})
        )
      )
    );
  }
});

var SearchHome = React.createClass({displayName: "SearchHome",
  componentDidMount: function() {
    $(".create-search-result").removeClass("hidden");
  },
  render: function() {
    $(window).off("popstate").on("popstate", function() {
      window.location = window.location.href;
    });

    return (
      React.createElement("div", {className: "search-home"}, 
        React.createElement(Header, {smallTitle: true}), 
        React.createElement(SearchPanel, null)
      )
    );
  }
});

React.render(
  React.createElement(SearchHome, null),
  document.getElementById('content')
);