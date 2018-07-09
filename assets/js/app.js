let app = angular.module("dbBackup", ["ngRoute"]);
app.config(function ($routeProvider) {
    $routeProvider
        .when("/", {
            templateUrl: 'assets/templates/home.html'
        })
        .when("/backups", {
            templateUrl: 'assets/templates/backups/backups.html',
            controller: 'backups'
        })
        .when("/backups/add", {
            templateUrl: 'assets/templates/backups/add.html',
            controller: 'addBackups'
        })
        .when("/schedule", {
            templateUrl: 'assets/templates/schedule.html',
            controller: 'schedule'
        })
        .when("/my-admin", {
            templateUrl: 'assets/templates/my-admin.html',
            controller: 'my_admin'
        })
        .when("/account", {
            templateUrl: 'assets/templates/account.html',
            controller: 'account'
        })
}).controller('backups', function ($scope, $http) {

}).controller('addBackups', function ($scope, $http) {

    $scope.submit = function () {
        $http('/', $scope.d).success(function (res) {

        })
    };

}).controller('schedule', function ($scope, $http) {

}).controller('my_admin', function ($scope, $http) {

}).controller('account', function ($scope, $http) {

});