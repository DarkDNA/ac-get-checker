This is a Go library that validates a ac-get repo for correctness.

Please do not use the github.com import path, but instead use http://ac-get.darkdna.net/checker instead.

Current Tests
=============

  * /desc.txt exists
  * /packages.list exists & is valid
  * packages have:
    * details.pkg exists & some basic recommendations.
    * All files mentioned in details.pkg exist.