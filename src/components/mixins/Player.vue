<script>
  var _ = require('lodash')

  export default {
    methods: {
      playerByPID: function (players, pid) {
        return _.find(players, ['id', pid])
      },
      playerByUID: function (players, uid) {
        return _.find(players, ['user.id', uid])
      },
      pidByUID: function (players, uid) {
        var self = this
        return _.get(self.playerByUID(players, uid), 'id', -1)
      },
      isPlayerFor: function (player, user) {
        var pid = _.get(player, 'user.id', -1)
        var uid = _.get(user, 'id', -2)
        var isAdmin = _.get(user, 'isAdmin', false)

        return isAdmin || pid === uid
      },
      currentPlayer: function (header, players) {
        var self = this
        var cpid = _.get(header.cpUserIndices, 0, -1)
        return self.playerByPID(players, cpid)
      },
      isCP: function (header, players, cu) {
        var self = this
        var cp = self.currentPlayer(header, players)
        return self.isPlayerFor(cp, cu)
      },
      isCPorAdmin: function (header, players, cu) {
        var self = this
        return cu.isAdmin || self.isCP(header, players, cu)
      }
    }
  }
</script>
