<template>
  <div id="landingContainer">
    <div id="upper">
      <div id="left">
        <div id="image_seal_container">
          <img id="image_seal" src="../assets/images/SealCircle.png" />
          <div id="updateAvailableTooltip">Update Available</div>
        </div>
      </div>
      <div id="content"></div>
      <div id="right">
        <div id="rightContainer">
          <div id="user_content">
            <span id="user_text">Username</span>
            <div id="avatarContainer">
              <button id="avatarOverlay">Edit</button>
            </div>
          </div>
          <div id="mediaContent">
            <div id="internalMedia">
              <div class="mediaContainer" id="settingsMediaContainer">
                <button class="mediaButton" id="settingsMediaButton">
                  <svg id="settingsSVG" class="mediaSVG" viewBox="0 0 141.36 137.43">
                    <path
                      d="M70.70475616319865,83.36934004916053 a15.320781354859122,15.320781354859122 0 1 1 14.454501310561755,-15.296030496450625 A14.850515045097694,14.850515045097694 0 0 1 70.70475616319865,83.36934004916053 M123.25082856443602,55.425620905968366 h-12.375429204248078 A45.54157947163293,45.54157947163293 0 0 0 107.21227231573047,46.243052436416285 l8.613298726156664,-9.108315894326587 a9.727087354538993,9.727087354538993 0 0 0 0,-13.167456673319956 l-3.465120177189462,-3.6631270444574313 a8.489544434114185,8.489544434114185 0 0 0 -12.375429204248078,0 l-8.613298726156664,9.108315894326587 A40.442902639482725,40.442902639482725 0 0 0 81.99114759747292,25.427580514871032 V12.532383284044531 a9.108315894326587,9.108315894326587 0 0 0 -8.811305593424633,-9.306322761594556 h-4.950171681699231 a9.108315894326587,9.108315894326587 0 0 0 -8.811305593424633,9.306322761594556 v12.895197230826497 a40.17064319698927,40.17064319698927 0 0 0 -9.331073620003052,4.0591407789933704 l-8.613298726156664,-9.108315894326587 a8.489544434114185,8.489544434114185 0 0 0 -12.375429204248078,0 L25.58394128451018,23.967279868769744 a9.727087354538993,9.727087354538993 0 0 0 0,13.167456673319956 L34.19724001066683,46.243052436416285 a45.07131316187151,45.07131316187151 0 0 0 -3.6631270444574313,9.083565035918088 h-12.375429204248078 a9.083565035918088,9.083565035918088 0 0 0 -8.811305593424633,9.306322761594556 v5.197680265784193 a9.108315894326587,9.108315894326587 0 0 0 8.811305593424633,9.306322761594556 h11.979415469712139 a45.69008462208391,45.69008462208391 0 0 0 4.0591407789933704,10.642869115653347 l-8.613298726156664,9.108315894326587 a9.727087354538993,9.727087354538993 0 0 0 0,13.167456673319956 l3.465120177189462,3.6631270444574313 a8.489544434114185,8.489544434114185 0 0 0 12.375429204248078,0 l8.613298726156664,-9.108315894326587 a40.49240435629971,40.49240435629971 0 0 0 9.331073620003052,4.0591407789933704 v12.895197230826497 a9.083565035918088,9.083565035918088 0 0 0 8.811305593424633,9.306322761594556 h4.950171681699231 A9.083565035918088,9.083565035918088 0 0 0 81.99114759747292,123.68848839660077 V110.79329116577425 a40.78941465720167,40.78941465720167 0 0 0 9.331073620003052,-4.0591407789933704 l8.613298726156664,9.108315894326587 a8.489544434114185,8.489544434114185 0 0 0 12.375429204248078,0 l3.465120177189462,-3.6631270444574313 a9.727087354538993,9.727087354538993 0 0 0 0,-13.167456673319956 l-8.613298726156664,-9.108315894326587 a45.665333763675406,45.665333763675406 0 0 0 4.034389920584874,-10.642869115653347 h12.004166328120636 a9.108315894326587,9.108315894326587 0 0 0 8.811305593424633,-9.306322761594556 v-5.197680265784193 a9.083565035918088,9.083565035918088 0 0 0 -8.811305593424633,-9.306322761594556 "
                      id="svg_3"
                      class
                    />
                  </svg>
                  <div id="settingsTooltip">Settings</div>
                </button>
              </div>
            </div>
            <div class="mediaDivider"></div>
            <div id="externalMedia">
              <div class="mediaContainer">
                <a href="https://www.northernblade.ru" class="mediaURL" id="linkURL">
                  <svg id="linkSVG" class="mediaSVG" viewBox="35.34 34.3575 70.68 68.71500">
                    <g>
                      <path
                        d="M75.37,65.51a3.85,3.85,0,0,0-1.73.42,8.22,8.22,0,0,1,.94,3.76A8.36,8.36,0,0,1,66.23,78H46.37a8.35,8.35,0,1,1,0-16.7h9.18a21.51,21.51,0,0,1,6.65-8.72H46.37a17.07,17.07,0,1,0,0,34.15H66.23A17,17,0,0,0,82.77,65.51Z"
                      />
                      <path
                        d="M66,73.88a3.85,3.85,0,0,0,1.73-.42,8.22,8.22,0,0,1-.94-3.76,8.36,8.36,0,0,1,8.35-8.35H95A8.35,8.35,0,1,1,95,78H85.8a21.51,21.51,0,0,1-6.65,8.72H95a17.07,17.07,0,0,0,0-34.15H75.13A17,17,0,0,0,58.59,73.88Z"
                      />
                    </g>
                  </svg>
                </a>
              </div>
              <div class="mediaContainer" style="display: none;">
                <a href="#" class="mediaURL" id="twitterURL" disabled>
                  <svg
                    id="twitterSVG"
                    class="mediaSVG"
                    viewBox="0 0 5000 4060"
                    preserveAspectRatio="xMidYMid meet"
                  >
                    <g>
                      <path
                        d="M1210 4048 c-350 -30 -780 -175 -1124 -378 -56 -33 -86 -57 -86 -68 0 -16 7 -17 83 -9 114 12 349 1 493 -22 295 -49 620 -180 843 -341 l54 -38 -49 -7 c-367 -49 -660 -256 -821 -582 -30 -61 -53 -120 -51 -130 3 -16 12 -17 73 -13 97 7 199 5 270 -4 l60 -9 -65 -22 c-341 -117 -609 -419 -681 -769 -18 -88 -26 -226 -13 -239 4 -3 32 7 63 22 68 35 198 77 266 86 28 4 58 9 68 12 10 2 -22 -34 -72 -82 -240 -232 -353 -532 -321 -852 15 -149 79 -347 133 -418 16 -20 17 -19 49 20 377 455 913 795 1491 945 160 41 346 74 485 86 l82 7 -7 -59 c-5 -33 -7 -117 -6 -189 2 -163 31 -286 103 -430 141 -285 422 -504 708 -550 112 -19 333 -19 442 0 180 30 335 108 477 239 l58 54 95 -24 c143 -36 286 -89 427 -160 70 -35 131 -60 135 -56 19 19 -74 209 -151 312 -50 66 -161 178 -216 217 l-30 22 73 -14 c111 -21 257 -63 353 -101 99 -39 99 -39 99 -19 0 57 -237 326 -412 468 l-88 71 6 51 c4 28 1 130 -5 226 -30 440 -131 806 -333 1202 -380 745 -1036 1277 -1823 1477 -243 62 -430 81 -786 78 -134 0 -291 -5 -349 -10z"
                      />
                    </g>
                  </svg>
                </a>
              </div>
              <div class="mediaContainer" style="display: none;">
                <a href="#" class="mediaURL" id="instagramURL" disabled>
                  <svg id="instagramSVG" class="mediaSVG" viewBox="0 0 5040 5040">
                    <defs>
                      <radialGradient id="instaFill" cx="30%" cy="107%" r="150%">
                        <stop offset="0%" stop-color="#fdf497" />
                        <stop offset="5%" stop-color="#fdf497" />
                        <stop offset="45%" stop-color="#fd5949" />
                        <stop offset="60%" stop-color="#d6249f" />
                        <stop offset="90%" stop-color="#285AEB" />
                      </radialGradient>
                    </defs>
                    <g>
                      <path
                        d="M1390 5024 c-163 -9 -239 -19 -315 -38 -281 -70 -477 -177 -660 -361 -184 -184 -292 -380 -361 -660 -43 -171 -53 -456 -53 -1445 0 -989 10 -1274 53 -1445 69 -280 177 -476 361 -660 184 -184 380 -292 660 -361 171 -43 456 -53 1445 -53 989 0 1274 10 1445 53 280 69 476 177 660 361 184 184 292 380 361 660 43 171 53 456 53 1445 0 989 -10 1274 -53 1445 -69 280 -177 476 -361 660 -184 184 -380 292 -660 361 -174 44 -454 53 -1470 52 -599 0 -960 -5 -1105 -14z m2230 -473 c58 -6 141 -18 185 -27 397 -78 638 -318 719 -714 37 -183 41 -309 41 -1290 0 -981 -4 -1107 -41 -1290 -81 -395 -319 -633 -714 -714 -183 -37 -309 -41 -1290 -41 -981 0 -1107 4 -1290 41 -397 81 -636 322 -714 719 -33 166 -38 296 -43 1100 -5 796 3 1203 27 1380 67 489 338 758 830 825 47 7 162 15 255 20 250 12 1907 4 2035 -9z"
                      />
                      <path
                        d="M2355 3819 c-307 -42 -561 -172 -780 -400 -244 -253 -359 -543 -359 -899 0 -361 116 -648 367 -907 262 -269 563 -397 937 -397 374 0 675 128 937 397 251 259 367 546 367 907 0 361 -116 648 -367 907 -197 203 -422 326 -690 378 -101 20 -317 27 -412 14z m400 -509 c275 -88 470 -284 557 -560 20 -65 23 -95 23 -230 0 -135 -3 -165 -23 -230 -88 -278 -284 -474 -562 -562 -65 -20 -95 -23 -230 -23 -135 0 -165 3 -230 23 -278 88 -474 284 -562 562 -20 65 -23 95 -23 230 0 135 3 165 23 230 73 230 219 403 427 507 134 67 212 83 390 79 111 -3 155 -8 210 -26z"
                      />
                      <path
                        d="M3750 1473 c-29 -11 -66 -38 -106 -77 -70 -71 -94 -126 -94 -221 0 -95 24 -150 94 -221 72 -71 126 -94 225 -94 168 0 311 143 311 311 0 99 -23 154 -94 225 -43 42 -76 66 -110 77 -61 21 -166 21 -226 0z"
                      />
                    </g>
                  </svg>
                </a>
              </div>
              <div class="mediaContainer" style="display: none;">
                <a href="#" class="mediaURL" id="youtubeURL" disabled>
                  <svg id="youtubeSVG" class="mediaSVG" viewBox="35.34 34.3575 70.68 68.71500">
                    <g>
                      <path
                        d="M84.8,69.52,65.88,79.76V59.27Zm23.65.59c0-5.14-.79-17.63-3.94-20.57S99,45.86,73.37,45.86s-28,.73-31.14,3.68S38.29,65,38.29,70.11s.79,17.63,3.94,20.57,5.52,3.68,31.14,3.68,28-.74,31.14-3.68,3.94-15.42,3.94-20.57"
                      />
                    </g>
                  </svg>
                </a>
              </div>
              <div class="mediaContainer">
                <a href="https://discord.gg/Px8yNer" class="mediaURL" id="discordURL">
                  <svg id="discordSVG" class="mediaSVG" viewBox="35.34 34.3575 70.68 68.71500">
                    <g>
                      <path
                        d="M81.23,78.48a6.14,6.14,0,1,1,6.14-6.14,6.14,6.14,0,0,1-6.14,6.14M60,78.48a6.14,6.14,0,1,1,6.14-6.14A6.14,6.14,0,0,1,60,78.48M104.41,73c-.92-7.7-8.24-22.9-8.24-22.9A43,43,0,0,0,88,45.59a17.88,17.88,0,0,0-8.38-1.27l-.13,1.06a23.52,23.52,0,0,1,5.8,1.95,87.59,87.59,0,0,1,8.17,4.87s-10.32-5.63-22.27-5.63a51.32,51.32,0,0,0-23.2,5.63,87.84,87.84,0,0,1,8.17-4.87,23.57,23.57,0,0,1,5.8-1.95l-.13-1.06a17.88,17.88,0,0,0-8.38,1.27,42.84,42.84,0,0,0-8.21,4.56S37.87,65.35,37,73s-.37,11.54-.37,11.54,4.22,5.68,9.9,7.14,7.7,1.47,7.7,1.47l3.75-4.68a21.22,21.22,0,0,1-4.65-2A24.47,24.47,0,0,1,47.93,82S61.16,88.4,70.68,88.4c10,0,22.75-6.44,22.75-6.44a24.56,24.56,0,0,1-5.35,4.56,21.22,21.22,0,0,1-4.65,2l3.75,4.68s2,0,7.7-1.47,9.89-7.14,9.89-7.14.55-3.85-.37-11.54"
                      />
                    </g>
                  </svg>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div id="lower">
      <div id="left">
        <div class="bot_wrapper" style="display: none;">
          <div id="content">
            <div id="server_status_wrapper">
              <span class="bot_label" id="landingPlayerLabel">SERVER</span>
              <span id="player_count">OFFLINE</span>
            </div>
            <div class="bot_divider"></div>
            <div id="mojangStatusWrapper">
              <span class="bot_label" id="landingBotLabel">MOJANG STATUS</span>
              <span id="mojang_status_icon">&#8226;</span>
              <div id="mojangStatusTooltip">
                <div id="mojangStatusTooltipTitle">Services</div>
                <div id="mojangStatusEssentialContainer">
                  <!-- Essential Mojang services are populated here. -->
                </div>
                <div id="mojangStatusNEContainer">
                  <div class="mojangStatusNEBar"></div>
                  <div id="mojangStatusNETitle">Non&nbsp;Essential</div>
                  <div class="mojangStatusNEBar"></div>
                </div>
                <div id="mojangStatusNonEssentialContainer">
                  <!-- Non Essential Mojang services are populated here. -->
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div id="center">
        <div class="bot_wrapper">
          <div id="content" style="display: none;">
            <button id="newsButton">
              <!--<img src="assets/images/icons/arrow.svg" id="newsButtonSVG"/>-->
              <div id="newsButtonAlert" style="display: none;"></div>
              <svg id="newsButtonSVG" viewBox="0 0 24.87 13.97">
                <defs>
                  <!-- <style>
  .arrowLine {
    fill: none;
    stroke: #fff;
    stroke-width: 2px;
  }
                  </style>-->
                </defs>
                <polyline class="arrowLine" points="0.71 13.26 12.56 1.41 24.16 13.02" />
              </svg>
              &#10;
              <span id="newsButtonText">NEWS</span>
            </button>
          </div>
        </div>
      </div>
      <div id="right">
        <div class="bot_wrapper">
          <div id="launch_content">
            <button id="launch_button" v-on:click="launchGame">Играть</button>
            <div class="bot_divider"></div>
            <button
              id="server_selection_button"
              class="bot_label"
            >&#8226; Northern Blade Main Server</button>
          </div>
          <div id="launch_details">
            <div id="launch_details_left">
              <span id="launch_progress_label">0%</span>
              <div class="bot_divider"></div>
            </div>
            <div id="launch_details_right">
              <progress id="launch_progress" value="22" max="100"></progress>
              <span id="launch_details_text" class="bot_label">Please wait..</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div id="newsContainer">
      <div id="newsContent" article="-1" style="display: none;">
        <div id="newsStatusContainer">
          <div id="newsStatusContent">
            <div id="newsTitleContainer">
              <a id="newsArticleTitle" href="#">Lorem Ipsum</a>
            </div>
            <div id="newsMetaContainer">
              <div id="newsArticleDateWrapper">
                <span id="newsArticleDate">Mar 15, 44 BC, 9:14 AM</span>
              </div>
              <div id="newsArticleAuthorWrapper">
                <span id="newsArticleAuthor">by Cicero</span>
              </div>
              <a href="#" id="newsArticleComments">0 Comments</a>
            </div>
          </div>
          <div id="newsNavigationContainer">
            <button id="newsNavigateLeft">
              <svg id="newsNavigationLeftSVG" viewBox="0 0 24.87 13.97">
                <defs>
                  <!-- <style>
  .arrowLine {
    fill: none;
    stroke: #fff;
    stroke-width: 2px;
    transition: 0.25s ease;
  }
                  </style>-->
                </defs>
                <polyline class="arrowLine" points="0.71 13.26 12.56 1.41 24.16 13.02" />
              </svg>
            </button>
            <span id="newsNavigationStatus">1 of 1</span>
            <button id="newsNavigateRight">
              <svg id="newsNavigationRightSVG" viewBox="0 0 24.87 13.97">
                <defs>
                  <!-- <style>
  .arrowLine {
    fill: none;
    stroke: #fff;
    stroke-width: 2px;
    transition: 0.25s ease;
  }
                  </style>-->
                </defs>
                <polyline class="arrowLine" points="0.71 13.26 12.56 1.41 24.16 13.02" />
              </svg>
            </button>
          </div>
        </div>
        <div id="newsArticleContainer">
          <div id="newsArticleContent">
            <div id="newsArticleContentScrollable">
              <!-- Article Content -->
            </div>
          </div>
        </div>
      </div>
      <div id="newsErrorContainer">
        <div id="newsErrorLoading">
          <span id="nELoadSpan" class="newsErrorContent">Checking for News..</span>
        </div>
        <div id="newsErrorFailed" style="display: none;">
          <span id="nEFailedSpan" class="newsErrorContent">Failed to Load News</span>
          <button id="newsErrorRetry">Try Again</button>
        </div>
        <div id="newsErrorNone" style="display: none;">
          <span id="nENoneSpan" class="newsErrorContent">No News</span>
        </div>
      </div>
    </div>
    <!--<script src="./assets/js/scripts/landing.js"></script> --->
  </div>
</template>

<script>
export default {
  data() {
    return 
  },

  methods: {
    launchGame() {
      window.backend.LauncherApplication.RunGame()

    }

  }

}
</script>

<style scoped>
/* Github Code Highlighting. */
/* @import "../../../node_modules/github-syntax-dark/lib/github-dark.css";
@import "../fonts/AvenirNextCyr/stylesheet.css"; */

/*******************************************************************************
 *                                                                             *
 * Fonts                                                                       *
 *                                                                             *
 ******************************************************************************/

@font-face {
  font-family: "Avenir Book";
  src: url("/../fonts/Avenir-Book.ttf");
}

@font-face {
  font-family: "Avenir Medium";
  src: url("/../fonts/Avenir-Medium.ttf");
}

@font-face {
  font-family: "Ringbearer";
  src: url("/../fonts/Ringbearer.ttf");
}

/*******************************************************************************
 *                                                                             *
 * Element Styles                                                              *
 *                                                                             *
 ******************************************************************************/

/* Reset body, html, and div presets. */
body,
html,
div {
  margin: 0px;
  padding: 0px;
}

/* Reset p presets. */
p {
  -webkit-margin-before: 0em;
  -webkit-margin-after: 0em;
}

/* Set default font and color. */
body,
button {
  font-family: "Avenir Book";
  color: white;
}

body {
  background: url("/../images/backgrounds/2.jpg") no-repeat center center fixed;
  background-size: cover;
}

/*******************************************************************************
 *                                                                             *
 * Frame Styles (frame.ejs)                                                    *
 *                                                                             *
 ******************************************************************************/

/* Frame Bar */
#frameBar {
  position: relative;
  z-index: 100;
  display: flex;
  flex-direction: column;
  transition: background-color 1s ease;
  /*background-color: rgba(0, 0, 0, 0.5);*/
  -webkit-user-select: none;
}

/* Undraggable region on the top of the frame. */
#frameResizableTop {
  height: 2px;
  width: 100%;
  -webkit-app-region: no-drag;
}

/* Flexbox to wrap the main frame content. */
#frameMain {
  display: flex;
  height: 20px;
}

/* Undraggable region on the left and right of the frame. */
.frameResizableVert {
  width: 2px;
  -webkit-app-region: no-drag;
}

/* Main frame content for windows. */
#frameContentWin {
  display: flex;
  justify-content: space-between;
  width: 100%;
  -webkit-app-region: drag;
}

/* Main frame content for darwin. */
#frameContentDarwin {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  width: 100%;
  -webkit-app-region: drag;
}

/* Frame logo (windows only). */
#frameTitleDock {
  padding: 0px 10px;
}
#frameTitleText {
  font-size: 14px;
  font-family: "Avenir Medium";
  letter-spacing: 0.5px;
}

/* Windows frame button dock. */
#frameButtonDockWin {
  -webkit-app-region: no-drag !important;
  position: relative;
  top: -2px;
  right: -2px;
  height: 22px;
}
#frameButtonDockWin > .frameButton:not(:first-child) {
  margin-left: -4px;
}

/* Darwin frame button dock: NaN; */
#frameButtonDockDarwin {
  -webkit-app-region: no-drag !important;
  position: relative;
  top: -1px;
  right: -1px;
}

/* Windows Frame Button Styles. */
.frameButton {
  background: none;
  border: none;
  height: 22px;
  width: 39px;
  cursor: pointer;
}
.frameButton:hover,
.frameButton:focus {
  background: rgba(189, 189, 189, 0.43);
}
.frameButton:active {
  background: rgba(156, 156, 156, 0.43);
}
.frameButton:focus {
  outline: 0px;
}

/* Close button is red. */
#frameButton_close:hover,
#frameButton_close:focus {
  background: rgba(255, 53, 53, 0.61) !important;
}
#frameButton_close:active {
  background: rgba(235, 0, 0, 0.61) !important;
}

/* Darwin Frame Button Styles. */
.frameButtonDarwin {
  height: 12px;
  width: 12px;
  border-radius: 50%;
  border: 0px;
  margin-left: 5px;
  -webkit-app-region: no-drag !important;
  cursor: pointer;
}
.frameButtonDarwin:focus {
  outline: 0px;
}

#frameButtonDarwin_close {
  background-color: #e74c32;
}
#frameButtonDarwin_close:hover,
#frameButtonDarwin_close:focus {
  background-color: #ff9a8a;
}
#frameButtonDarwin_close:active {
  background-color: #ff8d7b;
}

#frameButtonDarwin_minimize {
  background-color: #fed045;
}
#frameButtonDarwin_minimize:hover,
#frameButtonDarwin_minimize:focus {
  background-color: #ffe9a9;
}
#frameButtonDarwin_minimize:active {
  background-color: #ffde7b;
}

#frameButtonDarwin_restoredown {
  background-color: #96e734;
}
#frameButtonDarwin_restoredown:hover,
#frameButtonDarwin_restoredown:focus {
  background-color: #d6ffa6;
}
#frameButtonDarwin_restoredown:active {
  background-color: #bfff76;
}

/*******************************************************************************
 *                                                                             *
 * Welcome View (welcome.ejs)                                                  *
 *                                                                             *
 ******************************************************************************/

#welcomeContainer {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
}

#welcomeContent {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 50%;
  top: -10%;
  position: relative;
}

/*
.cloudDiv {
    position: absolute;
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
}

.cloudTop {
    height: 50%;
    width: 100%;
    background-image: url('../images/cloudTrans.png');
    animation: clouds1 80s linear infinite;
    background-size: cover;
}

.cloudBottom {
    height: 50%;
    width: 100%;
    background-image: url('../images/cloudTrans2.png');
    animation: clouds2 70s linear infinite;
    background-size: cover;
}

@keyframes clouds1 {
    to {
        background-position: 200%;
    }
}
@keyframes clouds2 {
    to {
        background-position: 230%;
    }
}
*/

#welcomeImageSeal {
  border-radius: 50%;
  border: 2px solid #cad7e1;
  background: rgba(1, 2, 1, 0.5);
  height: 125px;
  width: 125px;
  box-shadow: 0px 0px 10px 0px rgb(0, 0, 0);
  margin-bottom: 5%;
  margin-top: 10%;
}

#welcomeHeader {
  font-family: "Avenir Medium";
  text-align: center;
  color: white;
  margin-bottom: 25px;
  letter-spacing: 1px;
  font-size: 20px;
  text-shadow: white 0px 0px 0px;
}

#welcomeDescription {
  text-align: justify;
  font-size: 13px;
  font-weight: 100;
  text-shadow: rgba(255, 255, 255, 0.75) 0px 0px 20px;
}

#welcomeDescCTA {
  text-align: center;
  font-size: 14px;
  font-weight: 100;
  text-shadow: rgba(255, 255, 255, 0.75) 0px 0px 20px;
}

/* Login button styles. */
#welcomeButton {
  background: none;
  font-weight: bold;
  letter-spacing: 2px;
  border: none;
  padding: 15px 5px;
  margin: 10px 0px;
  cursor: pointer;
  position: relative;
  right: -20px;
  transition: 0.5s ease;
  margin-top: 5%;
  margin-bottom: -5%;
}
#welcomeButton:disabled {
  color: rgba(255, 255, 255, 0.75);
  pointer-events: none;
}
#welcomeButton:hover,
#welcomeButton:focus {
  text-shadow: 0px 0px 20px #fff;
  outline: none;
}
#welcomeButton:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7;
}
#welcomeSVG {
  -webkit-transform: translate3d(0, 0, 0);
  overflow: visible;
  transform: rotate(90deg);
  margin-left: 20px;
  transition: 0.25s ease;
  width: 20px;
  height: 20px;
}
#welcomeButton:hover #welcomeSVG,
#welcomeButton:focus #welcomeSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #fff);
}
#welcomeButton:active #welcomeSVG .arrowLine {
  stroke: #c7c7c7;
}
#welcomeButton:active #welcomeSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #c7c7c7);
}
#welcomeButton:disabled #welcomeSVG .arrowLine {
  stroke: rgba(255, 255, 255, 0.75);
}

#welcomeButtonContent {
  display: flex;
  align-items: center;
}

/*******************************************************************************
 *                                                                             *
 * Login View (login.ejs)                                                      *
 *                                                                             *
 ******************************************************************************/

/* Styles for dimmer login span. */
.loginSpanDim {
  font-size: 12px;
  color: #848484;
  font-weight: bold;
}

/* Main login container. */
#loginContainer {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  transition: filter 0.25s ease;
  background: rgba(0, 0, 0, 0.5);
}

/* Login cancel button styles. */
#loginCancelContainer {
  position: absolute;
  top: 5%;
  right: 5%;
}

/* Login cancel button styles. */
#loginCancelButton {
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
  transition: 0.25s ease;
}
#loginCancelButton:hover #loginCancelIcon,
#loginCancelButton:hover #loginCancelText,
#loginCancelButton:focus #loginCancelIcon,
#loginCancelButton:focus #loginCancelText {
  text-shadow: 0px 0px 20px white;
}
#loginCancelButton:hover #loginCancelIcon,
#loginCancelButton:focus #loginCancelIcon {
  box-shadow: 0px 0px 20px white;
}
#loginCancelButton:active #loginCancelIcon,
#loginCancelButton:active #loginCancelText {
  text-shadow: 0px 0px 20px rgba(255, 255, 255, 0.75);
  color: rgba(255, 255, 255, 0.75);
  border-color: rgba(255, 255, 255, 0.75);
}
#loginCancelButton:active #loginCancelIcon {
  box-shadow: 0px 0px 20px rgba(255, 255, 255, 0.75);
}
#loginCancelButton:disabled {
  pointer-events: none;
}
#loginCancelButton:disabled #loginCancelIcon,
#loginCancelButton:disabled #loginCancelText {
  color: rgba(255, 255, 255, 0.75);
  border-color: rgba(255, 255, 255, 0.75);
}

/* The X in a circle icon for the cancel button. */
#loginCancelIcon {
  border-radius: 50%;
  border: 1px solid white;
  box-sizing: border-box;
  height: 30px;
  width: 30px;
  font-size: 19px;
  line-height: 30px;
  margin: 0 auto;
  margin-bottom: 5px;
  transition: 0.25s ease;
}
/* Text for the login cancel button. */
#loginCancelText {
  font-size: 15px;
  transition: 0.25s ease;
}

/* Login content wrapper. */
#loginContent {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  padding: 0px 25px;
}

/* Login form. */
#loginForm {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* Login form anchor styles. */
#loginForm a {
  font-size: 12px;
  color: #848484;
  font-weight: bold;
  text-decoration: none;
  transition: 0.25s ease;
}
#loginForm a:hover,
#loginForm a:focus {
  color: #a2a2a2;
  outline: none;
}
#loginForm a:active {
  color: #8b8b8b;
}

/* Logo on login form. */
#loginImageSeal {
  border-radius: 50%;
  border: 2px solid #cad7e1;
  background: rgba(1, 2, 1, 0.5);
  height: 100px;
  width: 100px;
  box-shadow: 0px 0px 10px 0px rgb(0, 0, 0);
  margin-bottom: 20px;
}

/* Header on login view. */
#loginSubheader {
  font-family: "Avenir Medium";
  margin-bottom: 25px;
  font-size: 1em;
  letter-spacing: 1px;
  font-weight: bold;
}

/* Container to organize login field elements. */
.loginFieldContainer {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* SVG icons on the login view. */
.loginSVG {
  fill: #fff;
  height: 20px;
  width: 20px;
}

/* Span which displays errors related to login field content. */
.loginErrorSpan {
  font-family: "Avenir Medium";
  font-weight: bold;
  font-size: 0.7em;
  color: #ff1b0c;
  width: 100%;
  text-align: right;
  position: absolute;
  top: 7px;
  opacity: 0;
  transition: 0.25s ease;
}

.shake {
  animation: shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
}

@keyframes shake {
  10%,
  90% {
    transform: translate3d(-1px, 0, 0);
  }

  20%,
  80% {
    transform: translate3d(2px, 0, 0);
  }

  30%,
  50%,
  70% {
    transform: translate3d(-4px, 0, 0);
  }

  40%,
  60% {
    transform: translate3d(4px, 0, 0);
  }
}

/* Login text input styles. */
.loginField {
  font-family: "Avenir Book";
  background: none;
  border-width: 1.5px 0px 0px 0px;
  border-style: solid;
  /* width: 250px; */
  width: 300px;
  margin-bottom: 20px;
  border-color: #fff;
  color: rgba(255, 255, 255, 0.75);
  font-weight: bold;
  text-align: center;
  box-sizing: border-box;
  padding: 7.5px;
  font-size: 10px;
  letter-spacing: 1px;
}
.loginField:focus {
  outline: none;
}
.loginField:disabled {
  color: rgba(255, 255, 255, 0.5);
}
.loginField::-webkit-input-placeholder {
  color: rgba(255, 255, 255, 0.75);
  font-size: 10px;
  letter-spacing: 1px;
  text-align: center;
  font-weight: bold;
}
.loginField:focus::-webkit-input-placeholder {
  color: transparent;
}

/* Add spacing between password field and options bar. */
#labelPassword {
  margin-bottom: 13px;
}

/* Container which contains the forgot and remember options. */
#loginOptions {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

/* Remember option text. */
#loginRememberText {
  padding-right: 10px;
  transition: 0.25s ease;
}

/* Login button styles. */
#loginButton {
  background: none;
  font-weight: bold;
  letter-spacing: 2px;
  border: none;
  padding: 15px 5px;
  margin: 10px 0px;
  cursor: pointer;
  position: relative;
  transition: 0.5s ease;
}
#loginButton:disabled {
  color: rgba(255, 255, 255, 0.75);
  pointer-events: none;
}
#loginButton[loading] {
  color: #fff;
}
#loginButton:hover,
#loginButton:focus {
  text-shadow: 0px 0px 20px #fff;
  outline: none;
}
#loginButton:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7;
}
#loginSVG {
  -webkit-transform: translate3d(0, 0, 0);
  overflow: visible;
  transform: rotate(90deg);
  margin-left: 20px;
  transition: 0.25s ease;
  width: 20px;
  height: 20px;
}
#loginButton:hover #loginSVG,
#loginButton:focus #loginSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #fff);
}
#loginButton:active #loginSVG .arrowLine {
  stroke: #c7c7c7;
}
#loginButton:active #loginSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #c7c7c7);
}
#loginButton:disabled #loginSVG .arrowLine {
  stroke: rgba(255, 255, 255, 0.75);
}

#loginButtonContent {
  display: flex;
  align-items: center;
}

#loginButton .circle-loader,
#loginButton[loading] #loginSVG {
  display: none;
}
#loginButton[loading] .circle-loader,
#loginButton #loginSVG {
  display: initial;
}

.circle-loader {
  margin-left: 20px;
  border: 2px solid rgba(255, 255, 255, 0.5);
  border-left-color: #ffffff;
  animation-name: loader-spin;
  animation-duration: 1s;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
  position: relative;
  display: inline-block;
  vertical-align: top;
  border-radius: 50%;
  width: 16px;
  height: 16px;
}
.load-complete {
  animation: none;
  border-color: #ffffff;
  transition: border 500ms ease-out;
}
.checkmark {
  display: none;
}
.checkmark.draw:after {
  animation-duration: 800ms;
  animation-timing-function: ease;
  animation-name: checkmark;
  transform: scaleX(-1) rotate(135deg);
}
.checkmark:after {
  opacity: 1;
  height: 8px;
  width: 4px;
  transform-origin: left top;
  border-right: 2px solid #ffffff;
  border-top: 2px solid #ffffff;
  content: "";
  left: 2px;
  top: 8px;
  position: absolute;
}
@keyframes loader-spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
@keyframes checkmark {
  0% {
    height: 0;
    width: 0;
    opacity: 1;
  }
  20% {
    height: 0;
    width: 4px;
    opacity: 1;
  }
  40% {
    height: 8px;
    width: 4px;
    opacity: 1;
  }
  100% {
    height: 8px;
    width: 4px;
    opacity: 1;
  }
}

/*.spinningCircle {
    margin-left: 20px;
    height: 16px;
    width: 16px;
    border-radius: 50%;
    border: 2px solid rgba(255,255,255,0);
    border-top-color: #ffffff;
    border-right-color: #ffffff;
    border-left-color: rgba(255, 255, 255, 0.50);
    border-bottom-color: rgba(255, 255, 255, 0.50);
    animation: single2 4s infinite linear;
}

@keyframes single2 {
    0% {
        transform: rotate(0deg);
    }
    100% {
        transform: rotate(720deg);
    }
}*/

/* Disclaimer container. */
#loginDisclaimer {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* Add spacing between register anchor and disclaimer. */
#loginRegisterSpan {
  margin-bottom: 5px;
}

/* Disclaimer text styles. */
.loginDisclaimerText {
  font-size: 7px;
  color: #848484;
  font-weight: bold;
  text-align: center;
}

/* * *
* Login View | Custom Checkbox
* * */

/* Checkbox container. */
#checkmarkContainer {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  position: relative;
  cursor: pointer;
  font-size: 22px;
  -webkit-user-select: none;
}

/* Hide the default checkbox. */
#checkmarkContainer input {
  opacity: 0;
  cursor: pointer;
  position: absolute;
}

/* Create a custom checkbox. */
.loginCheckmark {
  position: relative;
  height: 10px;
  width: 10px;
  border: 1px solid #848484;
  border-radius: 1px;
  background: none;
  transition: 0.25s ease;
}
/* On hover and focus, add a grey border color. */
#checkmarkContainer:hover input ~ *,
#checkmarkContainer input:focus ~ * {
  color: #a2a2a2;
  border-color: #a2a2a2;
}
/* On keydown, darken the checkbox a bit. */
#checkmarkContainer input:active ~ *:not(#loginRememberText) {
  color: #8d8d8d;
  border-color: #8d8d8d;
}
#checkmarkContainer[disabled] {
  pointer-events: none;
}
/* For checked -> #checkmarkContainer input:checked ~ * */
/* Create the checkmark/indicator (hidden when not checked). */
.loginCheckmark:after {
  content: "";
  display: none;
}
/* Show the checkmark when checked. */
#checkmarkContainer input:checked ~ .loginCheckmark:after {
  display: block;
}
/* Style the checkmark/indicator. */
#checkmarkContainer .loginCheckmark:after {
  position: absolute;
  left: 3.5px;
  top: 0.5px;
  width: 2px;
  height: 6px;
  border: solid #a2a2a2;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

/*
#login_filter {
    height: calc(100% - 22px);
    width: 100%;
    z-index: 9000;
    position: absolute;
    filter: blur(8px) contrast(0.9) brightness(1.0);
    background: url('./../images/backgrounds/0.jpg') no-repeat center center fixed;
    transform: scale(1.2);
    background-size: cover;
}
*/

/*******************************************************************************
 *                                                                             *
 * Settings View (sttings.ejs)                                                 *
 *                                                                             *
 ******************************************************************************/

/* Main settings container. */
#settingsContainer {
  position: relative;
  height: 100%;
  display: flex;
  background-color: rgba(0, 0, 0, 0.5);
  transition: background-color 0.25s cubic-bezier(0.02, 0.01, 0.47, 1);
}

/* Drop shadow displayed when content is scrolled out of view. */
#settingsContainer:before {
  content: "";
  background: linear-gradient(rgba(0, 0, 0, 0.25), transparent);
  width: 100%;
  height: 5px;
  position: absolute;
  opacity: 0;
  transition: opacity 0.25s ease;
}
#settingsContainer[scrolled]:before {
  opacity: 1;
}

/* Left hand side of the settings UI, for navigation. */
#settingsContainerLeft {
  padding-top: 4%;
  height: 100%;
  width: 25%;
  box-sizing: border-box;
}

/* Settings navigation container. */
#settingsNavContainer {
  height: 100%;
  display: flex;
  flex-direction: column;
}

/* Navigation header styles. */
#settingsNavHeader {
  height: 15%;
  display: flex;
  justify-content: center;
}
#settingsNavHeaderText {
  font-size: 20px;
}

/* Navigation items outer container. */
#settingsNavItemsContainer {
  height: 85%;
  display: flex;
  justify-content: center;
  box-sizing: border-box;
}

/* Navigation items content container. */
#settingsNavItemsContent {
  height: 100%;
  display: flex;
  flex-direction: column;
  position: relative;
}

/* Navigation item shared styles. */
.settingsNavItem {
  background: none;
  border: none;
  text-align: left;
  margin: 5px 0px;
  padding: 0px 20px;
  color: grey;
  cursor: pointer;
  outline: none;
  transition: 0.25s ease;
}
.settingsNavItem:hover,
.settingsNavItem:focus {
  color: #c1c1c1;
  text-shadow: 0px 0px 20px #c1c1c1;
}
.settingsNavItem[selected] {
  cursor: default;
  color: white;
  text-shadow: none;
}

/* Div to add some space between nav items. */
.settingsNavSpacer {
  height: 25px;
}

/* Content container for the done button. */
#settingsNavContentBottom {
  position: absolute;
  top: 65%;
}

/* Settings navigational divider. */
.settingsNavDivider {
  width: 75%;
  height: 1px;
  background: rgba(126, 126, 126, 0.57);
  margin-left: auto;
  margin-bottom: 25px;
}

/* Settings done button styles. */
#settingsNavDone {
  background: none;
  border: none;
  text-align: left;
  margin: 5px 0px;
  padding: 0px 20px;
  color: white;
  cursor: pointer;
  outline: none;
  transition: 0.25s ease;
}
#settingsNavDone:hover,
#settingsNavDone:focus {
  text-shadow: 0px 0px 20px white, 0px 0px 20px white, 0px 0px 20px white;
}
#settingsNavDone:active {
  text-shadow: 0px 0px 20px rgba(255, 255, 255, 0.75),
    0px 0px 20px rgba(255, 255, 255, 0.75),
    0px 0px 20px rgba(255, 255, 255, 0.75);
  color: rgba(255, 255, 255, 0.75);
}
#settingsNavDone:disabled {
  color: rgba(255, 255, 255, 0.75);
  pointer-events: none;
}

/* Right hand side of the settings container, for tabs. */
#settingsContainerRight {
  height: 100%;
  width: 75%;
  box-sizing: border-box;
}

/* Settings tab shared styles. */
.settingsTab {
  width: 100%;
  height: 100%;
  overflow-y: auto;
}
.settingsTab::-webkit-scrollbar {
  width: 2px;
}
.settingsTab::-webkit-scrollbar-track {
  display: none;
}
.settingsTab::-webkit-scrollbar-thumb {
  border-radius: 10px;
  box-shadow: inset 0 0 10px rgba(255, 255, 255, 0.5);
}

/* Add spacing to the top of each settings tab. */
.settingsTab > *:first-child {
  margin-top: 5%;
}

/* Add spacing to the bottom of each settings tab. */
.settingsTab > *:last-child {
  margin-bottom: 20%;
}

/* Tab header shared styles. */
.settingsTabHeader {
  display: flex;
  flex-direction: column;
  margin-bottom: 20px;
}
.settingsTabHeaderText {
  font-size: 20px;
  font-family: "Avenir Medium";
}
.settingsTabHeaderDesc {
  font-size: 12px;
}

/* Remove spin button from number inputs. */
#settingsContainer input[type="number"]::-webkit-inner-spin-button {
  -webkit-appearance: none;
}

/* Default styles for text/number inputs. */
#settingsContainer input[type="number"],
#settingsContainer input[type="text"] {
  color: white;
  background: rgba(0, 0, 0, 0.25);
  border-radius: 3px;
  border: 1px solid rgba(126, 126, 126, 0.57);
  font-family: "Avenir Book";
  transition: 0.25s ease;
}
#settingsContainer input[type="number"]:focus,
#settingsContainer input[type="text"]:focus {
  outline: none;
  border-color: rgba(126, 126, 126, 0.87);
}
#settingsContainer input[type="number"][error] {
  border-color: rgb(255, 27, 12);
  background: rgba(236, 0, 0, 0.25);
  color: rgb(255, 27, 12);
}

/* Styles for a generic settings entry. */
.settingsFieldContainer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 0px;
  width: 75%;
  border-bottom: 1px solid rgba(255, 255, 255, 0.5);
}
.settingsFieldLeft {
  display: flex;
  flex-direction: column;
}
.settingsFieldTitle {
  font-size: 14px;
  font-family: "Avenir Medium";
  color: rgba(255, 255, 255, 0.95);
}
.settingsFieldDesc {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.95);
  margin-top: 5px;
}
.settingsDivider {
  height: 1px;
  width: 75%;
  background: rgba(255, 255, 255, 0.25);
}

/* Toggle Switch */
.toggleSwitch {
  position: relative;
  display: inline-block;
  width: 40px;
  height: 20px;
  border-radius: 50px;
  box-sizing: border-box;
}
.toggleSwitch input {
  display: none;
}
.toggleSwitchSlider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(255, 255, 255, 0.35);
  transition: 0.4s;
  border-radius: 50px;
  border: 1px solid rgba(126, 126, 126, 0.57);
}
.toggleSwitchSlider:before {
  position: absolute;
  content: "";
  height: 13px;
  width: 16px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.75);
  border-radius: 50px;
  transition: 0.4s;
}
input:checked + .toggleSwitchSlider {
  background-color: rgb(31, 140, 11);
  /* box-shadow: inset 2px 1px 20px black; */
  border: 1px solid rgb(31, 140, 11);
}
input:checked + .toggleSwitchSlider:before {
  transform: translateX(15px);
}

/* Range Slider styles. */
.rangeSlider {
  width: 35%;
  height: 5px;
  margin: 15px 0px;
  background: grey;
  border-radius: 3px;
  position: relative;
}
.rangeSliderBar {
  position: absolute;
  background: #8be88b;
  width: 50%;
  height: 5px;
  border-radius: 3px 0px 0px 3px;
  transition: background 0.25s ease;
}
.rangeSliderTrack {
  position: absolute;
  top: -7.5px;
  width: 7px;
  height: 20px;
  background: white;
  border-radius: 3px;
  left: 50%;
  cursor: ew-resize;
}

/* File selectors */

/* Main container for File selectors. */
.settingsFileSelContainer {
  display: flex;
  flex-direction: column;
  border-bottom: 1px solid rgba(255, 255, 255, 0.5);
  margin-bottom: 20px;
  margin-top: 20px;
  width: 75%;
}

/* File selector title. */
.settingsFileSelTitle {
  margin-bottom: 10px;
}

/* Wrapper container for the actionable elements. */
.settingsFileSelActions {
  display: flex;
  width: 90%;
}

/* File selector icon settings. */
.settingsFileSelIcon {
  display: flex;
  align-items: center;
  background: rgba(126, 126, 126, 0.57);
  border-radius: 3px 0px 0px 3px;
  padding: 5px;
  transition: 0.25s ease;
}
.settingsFileSelSVG {
  width: 20px;
  height: 20px;
  fill: white;
}

/* Disabled text field which stores the selected file path. */
.settingsFileSelVal {
  border-radius: 0px !important;
  width: 100%;
  padding: 5px 10px;
  font-size: 12px;
  height: 30px;
}

/* File input for file selection. */
.settingsFileSelSel {
  width: 0px;
  height: 0px;
  opacity: 0;
}
.settingsFileSelSel::-webkit-file-upload-button {
  display: none;
}

/* Wrapper label to add a custom style to the file input. */
.settingsFileSelLabel {
  border-left: 0px;
  border-radius: 0px 3px 3px 0px;
  font-size: 12px;
  padding: 0px 5px;
  cursor: pointer;
  display: flex;
  align-items: center;
  background: rgba(126, 126, 126, 0.57);
  transition: 0.25s ease;
  white-space: nowrap;
}
.settingsFileSelLabel:hover,
.settingsFileSelLabel:focus,
.settingsFileSelSel:focus ~ #settingsJavaExecLabel {
  text-shadow: 0px 0px 20px white;
}

/* Description for the file selector. */
.settingsFileSelDesc {
  font-size: 10px;
  margin: 20px 0px;
  color: lightgrey;
  width: 89%;
}
.settingsFileSelDesc strong {
  font-family: "Avenir Medium";
}

/* * *
* Settings View (Account Tab)
* * */

/* Add account button styles. */
#settingsAddAccount {
  background: rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(126, 126, 126, 0.57);
  border-radius: 3px;
  height: 50px;
  width: 75%;
  text-align: left;
  padding: 0px 50px;
  cursor: pointer;
  outline: none;
  transition: 0.25s ease;
}
#settingsAddAccount:hover,
#settingsAddAccount:focus {
  background: rgba(54, 54, 54, 0.25);
  text-shadow: 0px 0px 20px white;
}

/* Settings auth accounts header. */
#settingsCurrentAccountsHeader {
  margin: 20px 0px;
}

/* Auth account list container styles. */
#settingsCurrentAccounts {
  margin-bottom: 5%;
}
#settingsCurrentAccounts > .settingsAuthAccount:not(:last-child) {
  margin-bottom: 10px;
}
#settingsCurrentAccounts > .settingsAuthAccount:not(:first-child) {
  margin-top: 10px;
}

/* Auth account shared styles. */
.settingsAuthAccount {
  display: flex;
  width: 75%;
  background: rgba(0, 0, 0, 0.25);
  border-radius: 3px;
  border: 1px solid rgba(126, 126, 126, 0.57);
}

/* Left hand side of an auth account element, for the skin image. */
.settingsAuthAccountLeft {
  padding: 5px 5px 5px 20px;
}

/* Image of the auth account's skin. */
.settingsAuthAccountImage {
  height: 115px;
}

/* Right hand side of the auth account, for info + actions. */
.settingsAuthAccountRight {
  display: flex;
  width: 100%;
}

/* Account details container. */
.settingsAuthAccountDetails {
  display: flex;
  flex-direction: column;
  justify-content: center;
  margin-left: 20px;
  width: 100%;
}
.settingsAuthAccountDetails > *:not(:last-child) {
  margin-bottom: 20px;
}

/* Account detail element styles. */
.settingsAuthAccountDetailPane {
  display: flex;
  flex-direction: column;
}
.settingsAuthAccountDetailTitle {
  font-size: 12px;
  color: grey;
  font-weight: bold;
  font-family: "Avenir Medium";
}
.settingsAuthAccountDetailValue {
  font-size: 14px;
  -webkit-user-select: initial;
}

/* Account actions container. */
.settingsAuthAccountActions {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: flex-end;
  padding: 10px;
}

/* Account select button shared styles. */
.settingsAuthAccountSelect {
  opacity: 0;
  border: none;
  white-space: nowrap;
  background: none;
  font-family: "Avenir Medium";
  outline: none;
  transition: 0.25s ease;
}
.settingsAuthAccountSelect:hover:not([selected]),
.settingsAuthAccountSelect:focus:not([selected]) {
  text-shadow: 0px 0px 20px white, 0px 0px 20px white;
  cursor: pointer;
}
.settingsAuthAccount:hover .settingsAuthAccountSelect:not([selected]),
.settingsAuthAccountSelect[selected] {
  opacity: 1;
}
.settingsAuthAccountSelect[selected] {
  pointer-events: none;
}

/* Account logout button shared styles. */
.settingsAuthAccountLogOut {
  opacity: 0;
  border: 1px solid rgb(241, 55, 55);
  color: rgb(241, 55, 55);
  background: none;
  font-size: 12px;
  border-radius: 3px;
  font-family: "Avenir Medium";
  transition: 0.25s ease;
  cursor: pointer;
  outline: none;
}
.settingsAuthAccountLogOut:hover,
.settingsAuthAccountLogOut:focus {
  box-shadow: 0px 0px 20px rgb(241, 55, 55);
  background: rgba(241, 55, 55, 0.25);
}
.settingsAuthAccountLogOut:active {
  box-shadow: 0px 0px 20px rgb(185, 47, 47);
  background: rgba(185, 47, 47, 0.25);
  border: 1px solid rgb(185, 47, 47);
  color: rgb(185, 47, 47);
}
.settingsAuthAccount:hover .settingsAuthAccountLogOut {
  opacity: 1;
}

/* * *
* Settings View (Minecraft Tab)
* * */

/* Game resolution UI elements. */
#settingsGameResolutionContainer {
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.5);
  width: 75%;
}
#settingsGameResolutionContent {
  display: flex;
  align-items: center;
  padding-top: 10px;
}
#settingsGameResolutionCross {
  color: grey;
  padding: 0px 15px;
}
#settingsGameWidth,
#settingsGameHeight {
  padding: 7.5px 5px;
  width: 75px;
}

/* * *
* Settings View (Mods Tab)
* * */

/* Selected server content container */
#settingsSelServContainer {
  background: rgba(0, 0, 0, 0.25);
  width: 75%;
  border-radius: 3px;
  display: flex;
  justify-content: space-between;
  margin: 15px 0px;
}

/* Div which will be populated with the selected server's information. */
#settingsSelServContent {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding: 5px 0px;
}

/* Wrapper container for the switch server button. */
#settingsSwitchServerContainer {
  display: flex;
  align-items: center;
  padding: 15px;
}

/* Button to switch server configurations on the mods tab. */
#settingsSwitchServerButton {
  opacity: 0;
  border: 1px solid rgb(255, 255, 255);
  color: rgb(255, 255, 255);
  background: none;
  font-size: 12px;
  border-radius: 3px;
  font-family: "Avenir Medium";
  transition: 0.25s ease;
  cursor: pointer;
  outline: none;
}
#settingsSwitchServerButton:hover,
#settingsSwitchServerButton:focus {
  box-shadow: 0px 0px 20px rgb(255, 255, 255);
  background: rgba(255, 255, 255, 0.25);
}
#settingsSwitchServerButton:active {
  box-shadow: 0px 0px 20px rgb(187, 187, 187);
  background: rgba(187, 187, 187, 0.25);
  border: 1px solid rgb(187, 187, 187);
  color: rgb(187, 187, 187);
}
#settingsSelServContainer:hover #settingsSwitchServerButton {
  opacity: 1;
}

/* Main content container for the mod elements. */
#settingsModsContainer {
  width: 75%;
}

/* Mod sub-container header text. */
.settingsModsHeader {
  padding-bottom: 10px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.5);
  margin-bottom: 10px;
}

/* Mod elements sub-containers. */
#settingsReqModsContainer,
#settingsOptModsContainer,
#settingsDropinModsContainer {
  padding-bottom: 25px;
}

/* Main content containers for mod elements. */
#settingsReqModsContent,
#settingsOptModsContent,
#settingsDropinModsContent {
  font-size: 12px;
  background: rgba(0, 0, 0, 0.25);
  border-radius: 3px;
  color: white;
}

/* Mod elements. */
.settingsMod,
.settingsDropinMod {
  padding: 10px;
}
.settingsSubMod {
  padding: 10px 0px 10px 15px;
  margin-left: 20px;
  border-left: 1px solid rgba(255, 255, 255, 0.5);
}

/* Main content container for mod element information. */
.settingsModContent {
  display: flex;
  align-items: center;
  justify-content: space-between;
  transition: opacity 0.25s ease;
}

/* Wrapper container for the left side of a mod element. */
.settingsModMainWrapper {
  display: flex;
  align-items: center;
}

/* Mod enabled/disabled status. */
.settingsModStatus {
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background-color: #c32625;
  margin-right: 15px;
  transition: 0.25s ease;
}

/* Mod details container. */
.settingsModDetails {
  display: flex;
  flex-direction: column;
}

/* The version of the mod. */
.settingsModVersion {
  color: grey;
  font-size: 10px;
}

/* Disabled toggleswitch for required mods. */
.toggleSwitch[reqmod] {
  filter: grayscale(49%) brightness(60%);
  pointer-events: none;
}

/* Set the status color of an enabled mod. */
.settingsBaseMod[enabled]
  > .settingsModContent
  > .settingsModMainWrapper
  > .settingsModStatus {
  background-color: rgb(165, 195, 37);
}

/* Add opacity to submods of a disabled mod. */
.settingsBaseMod:not([enabled]) > .settingsSubModContainer .settingsModContent {
  opacity: 0.5;
}

/* Curve the left border for submods. */
.settingsSubModContainer > .settingsSubMod:first-child {
  border-top-left-radius: 10px;
}
.settingsSubModContainer > .settingsSubMod:last-child {
  border-bottom-left-radius: 10px;
}
.settingsSubModContainer > .settingsSubMod:only-child {
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
}

/* Wrapper container for all submods. */
.settingsSubModContainer {
  margin-top: 10px;
}

/* Button to open the mods folder for drop-in mods. */
#settingsDropinFileSystemButton {
  background: rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(126, 126, 126, 0.57);
  border-radius: 3px;
  height: 50px;
  width: 100%;
  text-align: left;
  padding: 0px 50px;
  cursor: pointer;
  outline: none;
  transition: 0.25s ease;
  margin-bottom: 10px;
}
#settingsDropinFileSystemButton:hover,
#settingsDropinFileSystemButton:focus,
#settingsDropinFileSystemButton[drag] {
  background: rgba(54, 54, 54, 0.25);
  text-shadow: 0px 0px 20px white;
}
/* Refresh instructions on the file system button. */
#settingsDropinRefreshNote {
  font-size: 10px;
  pointer-events: none;
}

/* Button to remove drop-in mods. */
.settingsDropinRemoveButton {
  background: none;
  border: none;
  font-size: 10px;
  text-align: left;
  padding: 0px;
  color: #c32625;
  font-weight: bold;
  cursor: pointer;
  outline: none;
  transition: 0.25s ease;
}
.settingsDropinRemoveButton:hover,
.settingsDropinRemoveButton:focus {
  text-shadow: 0px 0px 20px #c32625, 0px 0px 20px #c32625, 0px 0px 20px #c32625;
}
.settingsDropinRemoveButton:active {
  color: #9b1f1f;
  text-shadow: 0px 0px 20px #9b1f1f, 0px 0px 20px #9b1f1f, 0px 0px 20px #9b1f1f;
}

/* Shaderpack settings description. */
#settingsShaderpackDesc {
  font-size: 10px;
  margin: 10px 0px;
  color: lightgrey;
  font-weight: bold;
  width: 89%;
}

/* Wrapper container. */
#settingsShaderpackWrapper {
  display: flex;
}

/* Button to add shaderpacks. */
#settingsShaderpackButton {
  background: rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(126, 126, 126, 0.57);
  border-radius: 3px;
  cursor: pointer;
  outline: none;
  transition: 0.25s ease;
  font-size: 14px;
  padding: 6px 11px;
  margin-right: 5px;
}
#settingsShaderpackButton:hover,
#settingsShaderpackButton:focus,
#settingsShaderpackButton[drag] {
  background: rgba(54, 54, 54, 0.25);
  text-shadow: 0px 0px 20px white;
}

/* Main select container. */
.settingsSelectContainer {
  position: relative;
  width: 50%;
}

/* Div which displays the selected option. */
.settingsSelectSelected {
  border-radius: 3px;
  border-width: 1px;
  font-size: 14px;
  padding: 6px 16px;
}

/* Style the arrow inside the select element. */
.settingsSelectSelected:after {
  position: absolute;
  content: "";
  top: calc(50% - 3px);
  right: 10px;
  width: 0;
  height: 0;
  border: 6px solid transparent;
  border-color: rgba(126, 126, 126, 0.57) transparent transparent transparent;
}

/* Point the arrow upwards when the select box is open (active). */
.settingsSelectSelected.select-arrow-active:after {
  border-color: transparent transparent rgba(126, 126, 126, 0.57) transparent;
  top: 7px;
}
.settingsSelectSelected.select-arrow-active {
  border-radius: 3px 3px 0px 0px;
}

/* Options content container. */
.settingsSelectOptions {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  z-index: 99;
  max-height: 300%;
  overflow-y: scroll;
  border: 1px solid rgba(126, 126, 126, 0.57);
  border-top: none;
  border-radius: 0px 0px 3px 3px;
}
/* Hide the items when the select box is closed. */
.settingsSelectOptions[hidden] {
  display: none;
}
.settingsSelectOptions::-webkit-scrollbar {
  width: 2px;
}
.settingsSelectOptions::-webkit-scrollbar-track {
  display: none;
}
.settingsSelectOptions::-webkit-scrollbar-thumb {
  border-radius: 10px;
  box-shadow: inset 0 0 10px rgba(255, 255, 255, 0.5);
}

/* Shared styles between options and selection div. */
.settingsSelectOptions div,
.settingsSelectSelected {
  background: rgba(0, 0, 0, 0.25);
  border-style: solid;
  border-color: rgba(126, 126, 126, 0.57);
  color: #ffffff;
  cursor: pointer;
}
.settingsSelectOptions div {
  border-width: 0px 0px 1px 0px;
  font-size: 12px;
  padding: 4px 16px;
}
.settingsSelectOptions div:last-child {
  border-bottom: none;
}

/* Hover + selected styles. */
.settingsSelectOptions div:hover,
.settingsSelectOptions div[selected] {
  background-color: rgba(255, 255, 255, 0.25) !important;
}

/* * *
* Settings View (Java Tab)
* * */

/* Style links on the Java tab. */
#settingsTabJava a,
.settingsChangelogText a {
  color: rgba(202, 202, 202, 0.75);
  transition: 0.25s ease;
  outline: none;
}
#settingsTabJava a:hover,
#settingsTabJava a:focus,
.settingsChangelogText a:hover,
.settingsChangelogText a:focus {
  color: rgba(255, 255, 255, 0.75);
}
#settingsTabJava a:active,
.settingsChangelogText a:active {
  color: rgba(165, 165, 165, 0.75);
}

/* Main container for memory management. */
#settingsMemoryContainer {
  width: 75%;
  display: flex;
  flex-direction: column;
  border-bottom: 1px solid rgba(255, 255, 255, 0.5);
  margin-bottom: 20px;
}

/* Memory management title. */
#settingsMemoryTitle {
  margin-bottom: 10px;
  padding-bottom: 5px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.5);
}

/* Memory management content. */
#settingsMemoryContent {
  display: flex;
  justify-content: space-between;
  width: 100%;
}
#settingsMemoryContentLeft {
  width: 69%;
}
#settingsMemoryContentRight {
  display: flex;
  align-items: center;
  margin-right: 10%;
}

/* Header for memory sliders. */
.settingsMemoryHeader {
  font-size: 14px;
}

/* Wrapper container for a memory slider and label. */
.settingsMemoryActionContainer {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

/* Label which displays a memory slider's value. */
.settingsMemoryLabel {
  font-size: 14px;
  margin-right: 2%;
}

/* Range sliders for min and max memory settings. */
#settingsMaxRAMRange,
#settingsMinRAMRange {
  width: 85%;
}

/* Memory status elements. */
#settingsMemoryStatus {
  display: flex;
  flex-direction: column;
}
#settingsMemoryStatus > .settingsMemoryStatusContainer:not(:last-child) {
  margin-bottom: 50%;
}
.settingsMemoryStatusContainer {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.settingsMemoryStatusTitle {
  font-size: 12px;
  color: grey;
  font-weight: bold;
}
.settingsMemoryStatusValue {
  color: lightgrey;
  font-size: 16px;
}

/* Description for memory management. */
#settingsMemoryDesc {
  font-size: 10px;
  margin: 20px 0px;
  color: lightgrey;
  font-weight: bold;
}

/* Status text which displays details on the selected executable. */
#settingsJavaExecDetails {
  font-weight: bold;
  color: grey;
  font-size: 12px;
}

/* Main container for the JVM options setting. */
#settingsJVMOptsContainer {
  width: 75%;
}

/* JVM options title. */
#settingsJVMOptsTitle {
  margin-bottom: 10px;
}

/* Wrapper container for the actionable elements. */
#settingsJVMOptsContent {
  display: flex;
  width: 90%;
}

/* Text field to input the JVM options. */
#settingsJVMOptsVal {
  border-radius: 0px 3px 3px 0px !important;
  width: 100%;
  padding: 5px 10px;
  font-size: 12px;
}
#settingsJVMOptsContent:focus-within > .settingsJavaIcon {
  background: rgba(126, 126, 126, 0.87);
}

/* Description for the JVM options setting. */
#settingsJVMOptsDesc {
  font-size: 10px;
  margin: 20px 0px;
  color: lightgrey;
  font-weight: bold;
  width: 89%;
}

/* * *
* Settings View (Launcher Tab)
* * */

/* Tailored style for the data directory header. */
#settingsDataDirTitle {
  margin-bottom: 10px;
}

/* * *
* Settings View (About Tab)
* * */

/* Main about content container. */
#settingsAboutCurrentContainer {
  display: flex;
  flex-direction: column;
  background: rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(126, 126, 126, 0.57);
  border-radius: 3px;
  width: 75%;
  margin-bottom: 20px;
}

/* About content. */
#settingsAboutCurrentContent {
  display: flex;
  flex-direction: column;
  padding: 15px;
}

/* About header elements. */
#settingsAboutCurrentHeadline {
  display: flex;
  align-items: center;
  padding-bottom: 5px;
  border-bottom: 1px solid rgba(126, 126, 126, 0.57);
}
#settingsAboutLogo {
  width: 30px;
  height: 30px;
  padding: 5px;
}
#settingsAboutTitle {
  font-size: 23px;
  padding-left: 10px;
}

/* Current version container. */
#settingsAboutCurrentVersion {
  display: flex;
  align-items: center;
  padding-top: 10px;
}

/* Checkmark next to the version information. */
#settingsAboutCurrentVersionCheck {
  border-radius: 50%;
  background: #23aa23;
  text-align: center;
  font-weight: bold;
  margin: 11px 12px;
  color: white;
  height: 15px;
  width: 15px;
  font-size: 12px;
  line-height: 17px;
}

/* Current version details container. */
#settingsAboutCurrentVersionDetails {
  margin-left: 10px;
}

/* Release type text. */
#settingsAboutCurrentVersionTitle {
  font-size: 12px;
  font-family: "Avenir Medium";
  color: #23aa23;
  font-weight: bold;
}

/* Current version text. */
#settingsAboutCurrentVersionLine {
  font-size: 10px;
  color: grey;
  font-weight: bold;
}

/* About information links. */
#settingsAboutButtons {
  display: flex;
  padding: 0px 15px;
  margin-bottom: 5px;
}
.settingsAboutButton {
  background: none;
  border: none;
  font-size: 10px;
  color: grey;
  padding: 0px 5px;
  transition: 0.25s ease;
  outline: none;
  text-decoration: none;
}
.settingsAboutButton:hover,
.settingsAboutButton:focus {
  color: rgb(165, 165, 165);
}
.settingsAboutButton:active {
  color: rgba(124, 124, 124, 0.75);
}

/* Main changelog container. */
.settingsChangelogContainer {
  display: flex;
  flex-direction: column;
  background: rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(126, 126, 126, 0.57);
  border-radius: 3px;
  width: 75%;
  margin-bottom: 20px;
}

/* Changelog content container. */
.settingsChangelogContent {
  display: flex;
  flex-direction: column;
  padding: 15px;
}

/* Changelog header container. */
.settingsChangelogHeadline {
  padding-bottom: 10px;
  margin-bottom: 10px;
  border-bottom: 1px solid rgba(126, 126, 126, 0.57);
}
/* Changelog header label. */
.settingsChangelogLabel {
  font-size: 12px;
  color: grey;
  font-weight: bold;
}

/* Changelog text content container. */
.settingsChangelogText {
  font-size: 12px;
}

/* Styles for the changelog elements. */
.settingsChangelogText p {
  margin-bottom: 16px;
  line-height: 1.5;
}
.settingsChangelogText blockquote {
  border-left: 0.25em solid rgba(126, 126, 126, 0.95);
  margin: 0px;
  padding: 0 0 0 1em;
  color: rgba(255, 255, 255, 0.85);
}
.settingsChangelogText code {
  padding: 0.1em 0.4em;
  font-size: 85%;
  background-color: rgba(255, 255, 255, 0.25);
  color: white;
  border-radius: 3px;
  font-family: "Avenir Book";
}
.settingsChangelogText li + li {
  margin-top: 0.25em;
}
.settingsChangelogText a.commit-link {
  font-weight: 400;
  color: #ffffff;
  text-decoration: none;
}
.settingsChangelogText a.commit-link:hover {
  text-decoration: underline !important;
  text-decoration-color: black;
}
.settingsChangelogText tt {
  padding: 0.1em 0.4em;
  font-size: 86%;
  background-color: white;
  border-radius: 3px;
  color: black;
  font-weight: bold;
}
.settingsChangelogText a.commit-link:hover tt {
  text-decoration: underline;
  text-decoration-color: black;
}
.settingsChangelogText .highlight {
  background: rgba(0, 0, 0, 0.3);
  user-select: initial;
  padding: 5px 10px;
}
.settingsChangelogText .highlight pre {
  margin: 0px;
}

/* Container for the changelog button. */
.settingsChangelogActions {
  padding: 0px 15px 5px 15px;
}

/* Open changelog on GitHub. */
.settingsChangelogButton {
  padding: 0px;
}

/* * *
* Settings View (Updates Tab)
* * */

/* Main about content container. */
#settingsUpdateStatusContainer {
  display: flex;
  flex-direction: column;
  background: rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(126, 126, 126, 0.57);
  border-radius: 3px;
  width: 75%;
  margin-bottom: 20px;
}

/* Update content. */
#settingsUpdateStatusContent {
  display: flex;
  flex-direction: column;
  padding: 15px;
}

/* Update header elements. */
#settingsUpdateStatusHeadline {
  display: flex;
  align-items: center;
  padding-bottom: 5px;
  border-bottom: 1px solid rgba(126, 126, 126, 0.57);
}
#settingsUpdateTitle {
  font-size: 16px;
  padding-left: 10px;
  font-weight: bold;
}

/* Update version container. */
#settingsUpdateVersion {
  display: flex;
  align-items: center;
  padding: 10px 0px;
  border-bottom: 1px solid rgba(126, 126, 126, 0.57);
}

/* Checkmark next to the version information. */
#settingsUpdateVersionCheck {
  border-radius: 50%;
  background: #23aa23;
  text-align: center;
  font-weight: bold;
  margin: 11px 12px;
  color: white;
  height: 15px;
  width: 15px;
  font-size: 12px;
  line-height: 17px;
}

/* Update version details container. */
#settingsUpdateVersionDetails {
  margin-left: 10px;
}

/* Release type text. */
#settingsUpdateVersionTitle {
  font-size: 12px;
  font-family: "Avenir Medium";
  color: #23aa23;
  font-weight: bold;
}

/* Current version text. */
#settingsUpdateVersionLine {
  font-size: 10px;
  color: grey;
  font-weight: bold;
}

/* Update action container. */
#settingsUpdateActionContainer {
  padding-top: 10px;
  font-size: 14px;
  font-weight: bold;
}

/* Update action button styles. */
#settingsUpdateActionButton {
  display: flex;
  flex-direction: column;
  padding-left: 10px;
  background: none;
  border: none;
  font-size: 14px;
  font-weight: bold;
  cursor: pointer;
  outline: none;
  text-align: left;
  transition: 0.25s ease;
}
#settingsUpdateActionButton:hover,
#settingsUpdateActionButton:focus {
  text-shadow: 0px 0px 20px white, 0px 0px 20px white, 0px 0px 20px white;
}
#settingsUpdateActionButton:active {
  text-shadow: 0px 0px 20px #c7c7c7, 0px 0px 20px #c7c7c7, 0px 0px 20px #c7c7c7;
  color: #c7c7c7;
}
#settingsUpdateActionButton:disabled {
  pointer-events: none;
}

/*******************************************************************************
 *                                                                             *
 * Landing View (Structural Styles)                                            *
 *                                                                             *
 ******************************************************************************/

/* Main content container. */
#landingContainer {
  height: 100%;
  position: relative;
  transition: background 2s ease;
  overflow-y: hidden;
}

/* Upper content container. */
#landingContainer > #upper {
  position: relative;
  transition: top 2s ease;
  top: 0px;
  height: 77%;
  display: flex;
}
#landingContainer > #upper > #left {
  display: inline-flex;
  width: 15%;
  height: 100%;
  justify-content: flex-end;
}
#landingContainer > #upper > #content {
  display: inline-flex;
  width: 70%;
  height: 100%;
}
#landingContainer > #upper > #right {
  display: inline-flex;
  width: 15%;
  height: 100%;
}

/* Lower content container. */
#landingContainer > #lower {
  height: 23%;
  display: flex;
  background: linear-gradient(to top, rgba(0, 0, 0, 0.75), rgba(0, 0, 0, 0));
}
#landingContainer > #lower > #left {
  position: relative;
  transition: top 2s ease;
  top: 0px;
  height: 100%;
  width: 33%;
  display: inline-flex;
  justify-content: center;
}
#landingContainer > #lower > #left #content {
  position: relative;
  top: 25px;
  display: inline-flex;
  line-height: 24px;
  left: 50px;
}
#landingContainer > #lower > #center {
  position: relative;
  transition: top 2s ease;
  top: 0px;
  height: 100%;
  width: 34%;
  display: inline-flex;
  justify-content: center;
}
#landingContainer > #lower > #center #content {
  position: relative;
  z-index: 500;
  transition: top 2s ease;
  top: 10px;
}
#landingContainer > #lower > #right {
  position: relative;
  transition: top 2s ease;
  top: 0px;
  height: 100%;
  width: 33%;
  display: inline-flex;
}

/*******************************************************************************
 *                                                                             *
 * Landing View (News Styles)                                                  *
 *                                                                             *
 ******************************************************************************/

/* Main container. */
#newsContainer {
  position: absolute;
  top: 100%;
  height: 100%;
  width: 100%;
  transition: top 2s ease;
  display: flex;
  align-items: flex-end;
  justify-content: center;
}

/* News content container. */
#newsContent {
  height: 82vh;
  width: 100%;
  display: flex;
  -webkit-user-select: initial;
  position: relative;
}

/* Drop shadow displayed when content is scrolled out of view. */
#newsContent:before {
  content: "";
  background: linear-gradient(rgba(0, 0, 0, 0.25), transparent);
  width: 100%;
  height: 5px;
  position: absolute;
  opacity: 0;
  transition: opacity 0.25s ease;
}
#newsContent[scrolled]:before {
  opacity: 1;
}

/* News article status container (left). */
#newsStatusContainer {
  width: calc(30% - 60px);
  height: calc(100% - 30px);
  padding: 15px 15px 15px 45px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  position: relative;
}

/* News status content. */
#newsStatusContent {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}

/* News title wrapper. */
#newsTitleContainer {
  display: flex;
  max-width: 90%;
}

/* News article title styles. */
#newsArticleTitle {
  font-size: 18px;
  font-weight: bold;
  font-family: "Avenir Medium";
  color: white;
  text-decoration: none;
  transition: 0.25s ease;
  outline: none;
  text-align: right;
}
#newsArticleTitle:hover,
#newsArticleTitle:focus {
  text-shadow: 0px 0px 20px white;
}
#newsArticleTitle:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7;
}

/* News meta container. */
#newsMetaContainer {
  display: flex;
  flex-direction: column;
}

/* Date and author wrappers. */
#newsArticleDateWrapper,
#newsArticleAuthorWrapper {
  display: flex;
  justify-content: flex-end;
}

/* Date and author shared styles. */
#newsArticleDate,
#newsArticleAuthor {
  display: inline-block;
  font-size: 10px;
  padding: 0px 5px;
  font-weight: bold;
  border-radius: 2px;
}

/* Date styles. */
#newsArticleDate {
  background: white;
  color: black;
  margin-top: 5px;
}

/* Author styles. */
#newsArticleAuthor {
  background: #a02d2a;
}

/* News article comments styles. */
#newsArticleComments {
  margin-top: 5px;
  display: inline-block;
  font-size: 10px;
  color: #ffffff;
  text-decoration: none;
  transition: 0.25s ease;
  outline: none;
  text-align: right;
}
#newsArticleComments:focus,
#newsArticleComments:hover {
  color: #e0e0e0;
}
#newsArticleComments:active {
  color: #c7c7c7;
}

/* Article content container (right). */
#newsArticleContainer {
  width: calc(100% - 25px);
  height: 100%;
  margin: 0px 0px 0px 25px;
}

/* Article content styles. */
#newsArticleContentScrollable {
  font-size: 12px;
  overflow-y: scroll;
  height: 100%;
  padding: 0px 15px 0px 15px;
}
#newsArticleContentScrollable img,
#newsArticleContentScrollable iframe {
  max-width: 95%;
  display: block;
  margin: 0 auto;
}
#newsArticleContentScrollable a {
  color: rgba(202, 202, 202, 0.75);
  transition: 0.25s ease;
  outline: none;
}
#newsArticleContentScrollable a:hover,
#newsArticleContentScrollable a:focus {
  color: rgba(255, 255, 255, 0.75);
}
#newsArticleContentScrollable a:active {
  color: rgba(165, 165, 165, 0.75);
}
#newsArticleContentScrollable::-webkit-scrollbar {
  width: 2px;
}
#newsArticleContentScrollable::-webkit-scrollbar-track {
  display: none;
}
#newsArticleContentScrollable::-webkit-scrollbar-thumb {
  border-radius: 10px;
  box-shadow: inset 0 0 10px rgba(255, 255, 255, 0.5);
}
.bbCodeSpoilerButton {
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
  font-size: 16px;
  transition: 0.25s ease;
  width: 100%;
  border-bottom: 1px solid white;
  padding-bottom: 15px;
}
.bbCodeSpoilerButton:hover,
.bbCodeSpoilerButton:focus {
  text-shadow: 0px 0px 20px #ffffff, 0px 0px 20px #ffffff, 0px 0px 20px #ffffff;
}
.bbCodeSpoilerButton:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7, 0px 0px 20px #c7c7c7, 0px 0px 20px #c7c7c7;
}
.bbCodeSpoilerText {
  display: none;
  padding: 15px 0px;
  border-bottom: 1px solid white;
}

#newsArticleContentWrapper {
  width: 80%;
}

.newsArticleSpacerTop {
  height: 15px;
}

/* Div to add spacing at the end of a news article. */
.newsArticleSpacerBot {
  height: 30px;
}

/* News navigation container. */
#newsNavigationContainer {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 10px;
  -webkit-user-select: none;
  position: absolute;
  bottom: 15px;
  right: 0px;
}

/* Navigation status span. */
#newsNavigationStatus {
  font-size: 12px;
  margin: 0px 15px;
}

/* Left and right navigation button styles. */
#newsNavigateLeft,
#newsNavigateRight {
  background: none;
  border: none;
  outline: none;
  height: 20px;
  cursor: pointer;
}
#newsNavigateLeft:hover #newsNavigationLeftSVG,
#newsNavigateLeft:focus #newsNavigationLeftSVG,
#newsNavigateRight:hover #newsNavigationRightSVG,
#newsNavigateRight:focus #newsNavigationRightSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #fff);
}
#newsNavigateLeft:active #newsNavigationLeftSVG .arrowLine,
#newsNavigateRight:active #newsNavigationRightSVG .arrowLine {
  stroke: #c7c7c7;
}
#newsNavigateLeft:active #newsNavigationLeftSVG,
#newsNavigateRight:active #newsNavigationRightSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #c7c7c7);
}
#newsNavigateLeft:disabled #newsNavigationLeftSVG .arrowLine,
#newsNavigateRight:disabled #newsNavigationRightSVG .arrowLine {
  stroke: rgba(255, 255, 255, 0.75);
}
#newsNavigationLeftSVG {
  transform: rotate(-90deg);
  width: 15px;
}
#newsNavigationRightSVG {
  transform: rotate(90deg);
  width: 15px;
}

/* News error (message) container. */
#newsErrorContainer {
  height: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}
#newsErrorFailed {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}

/* News error content (message). */
.newsErrorContent {
  font-size: 20px;
}
#newsErrorLoading {
  display: flex;
  width: 168.92px;
}
#nELoadSpan {
  white-space: pre;
}
/* News error retry button styles. */
#newsErrorRetry {
  font-size: 12px;
  font-weight: bold;
  cursor: pointer;
  background: none;
  border: none;
  outline: none;
  transition: 0.25s ease;
}
#newsErrorRetry:focus,
#newsErrorRetry:hover {
  text-shadow: 0px 0px 20px white;
}
#newsErrorRetry:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7;
}

/*******************************************************************************
 *                                                                             *
 * Landing View (Top Styles)                                                   *
 *                                                                             *
 ******************************************************************************/

/* * *
* Landing View (Top Styles) | Left Content
* * */

/* Logo image. */
#image_seal {
  height: 70px;
  width: auto;
  position: relative;
  border: 2px solid white;
  box-sizing: border-box;
  border-radius: 50%;
}

/* Logo container styles. */
#image_seal_container {
  position: relative;
  height: 70px;
  width: 70px;
  border-radius: 50%;
  margin-top: 50px;
}

/* Logo container styles w/ update. */
#image_seal_container[update] {
  cursor: pointer;
}
#image_seal_container[update]:before,
#image_seal_container[update]:after {
  cursor: pointer;
  position: absolute;
  content: "";
  height: 100%;
  width: 100%;
  top: 0%;
  left: 0%;
  border-radius: 50%;
  box-shadow: 0 0 15px #43c628;
  animation: glow-grow 4s ease-out infinite;
  background: rgba(0, 0, 0, 0.15);
}
#image_seal_container[update]:before {
  animation-delay: 2s;
}

/* Update available tooltip styles. */
#updateAvailableTooltip {
  cursor: pointer;
  visibility: hidden;
  opacity: 0;
  width: 100px;
  height: 15px;
  background-color: rgb(0, 0, 0);
  color: #fff;
  text-align: center;
  border-radius: 4px;
  padding: 2px;
  position: absolute;
  z-index: 1;
  top: 115%;
  left: -17.5px;
  font-family: "Avenir Medium";
  font-size: 12px;
  transition: visibility 0s linear 0.25s, opacity 0.25s ease;
}
#updateAvailableTooltip::after {
  content: " ";
  position: absolute;
  left: 50%;
  bottom: 100%;
  margin-left: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: transparent transparent rgb(0, 0, 0) transparent;
}
#image_seal_container[update]:hover #updateAvailableTooltip {
  visibility: visible;
  opacity: 1;
  transition-delay: 0s;
}

/* Update available animation. */
@keyframes glow-grow {
  0% {
    opacity: 0;
    transform: scale(1);
  }
  80% {
    opacity: 1;
  }
  100% {
    transform: scale(1.5);
    opacity: 0;
  }
}

/* * *
* Landing View (Bottom Styles) | Right Content
* * */

/*  Wrapper container for top, right content. */
#rightContainer {
  display: flex;
  flex-direction: column;
  position: relative;
  top: 50px;
  align-items: flex-start;
  height: calc(100% - 50px);
}

/* Right hand user content container. */
#user_content {
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
  position: relative;
}

/* User profile avatar container. */
#avatarContainer {
  border-radius: 50%;
  border: 2px solid #cad7e1;
  box-sizing: border-box;
  background: rgba(1, 2, 1, 0.5);
  height: 70px;
  width: 70px;
  box-shadow: 0px 0px 10px 0px rgb(0, 0, 0);
  overflow: hidden;
  position: relative;
  background-position: center;
  background-repeat: no-repeat;
  background-size: contain;
}

/* Avatar edit overlay. */
#avatarOverlay {
  opacity: 0;
  position: absolute;
  z-index: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: 0.25s ease;
  font-weight: bold;
  letter-spacing: 2px;
  background-color: rgba(0, 0, 0, 0.35);
  -webkit-user-select: none;
  border: none;
  cursor: pointer;
  width: 100%;
  height: 100%;
  border-radius: 50%;
}
#avatarOverlay:hover,
#avatarOverlay:focus {
  opacity: 1;
}
#avatarOverlay:active {
  background-color: rgba(0, 0, 0, 0.45);
}

/* User profile name text. */
#user_text {
  font-size: 12px;
  min-width: 135px;
  font-weight: 900;
  letter-spacing: 1px;
  text-shadow: 0px 0px 20px black;
  position: absolute;
  right: 95px;
  text-align: right;
  -webkit-user-select: initial;
}

/* Social media icon content container. */
#mediaContent {
  position: relative;
  display: flex;
  flex-direction: column;
  margin-top: 25px;
  height: calc(100% - 95px);
  width: 70px;
  align-items: center;
}

/* Social Media Icon division containers. */
#internalMedia,
#externalMedia {
  display: flex;
  flex-direction: column;
}

/* Container object which wraps an icon to ensure fluid transitions. */
.mediaContainer {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 27px;
}

/* Divider bar between the external and internal icons. */
.mediaDivider {
  height: 1px;
  width: 14px;
  background: rgb(255, 255, 255);
  margin: 10px 0px;
}

/* Social media icon shared styles. */
.mediaSVG {
  fill: #ffffff;
  height: 12px;
  transition: 0.25s ease;
  cursor: pointer;
  height: 12px;
  width: 25px;
}
.mediaSVG:hover,
.mediaURL:focus .mediaSVG,
.mediaSVG:active {
  height: 20px;
}

/* Social media URL shared styles. */
.mediaURL {
  outline: none;
}

/* Internal media button shared styles. */
.mediaButton {
  background: none;
  border: none;
  padding: 0px;
  display: flex;
  align-items: center;
  outline: none;
}

#settingsMediaContainer {
  position: relative;
}

/* Settings icon colors. */
#settingsSVG {
  stroke: #ffffff;
  height: 15px;
}
.mediaButton:hover #settingsSVG,
.mediaButton:focus #settingsSVG,
.mediaButton:active #settingsSVG {
  height: 23px;
}

/* Settings tooltip styles. */
#settingsTooltip {
  visibility: hidden;
  opacity: 0;
  width: 75px;
  height: 20px;
  background-color: rgba(0, 0, 0, 0.75);
  text-align: center;
  border-radius: 4px;
  position: absolute;
  z-index: 1;
  right: 130%;
  font-size: 12px;
  line-height: 20px;
  transition: visibility 0s linear 0.25s, opacity 0.25s ease;
}
#settingsTooltip::after {
  content: " ";
  position: absolute;
  top: 50%;
  left: 100%;
  margin-top: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: transparent transparent transparent rgba(0, 0, 0, 0.75);
}
.mediaButton:hover #settingsTooltip,
.mediaButton:focus #settingsTooltip,
.mediaButton:active #settingsTooltip {
  visibility: visible;
  opacity: 1;
  transition-delay: 0s;
}

/* Twitter icon colors. */
#twitterSVG:hover,
#twitterURL:focus #twitterSVG {
  fill: #1da1f2;
}
#twitterSVG:active {
  fill: #1b8dd4;
}

/* Instagram icon colors. */
#instagramSVG:hover,
#instagramURL:focus #instagramSVG {
  fill: url("#instaFill");
  /*fill: radial-gradient(circle at 30% 107%, #fdf497 0%, #fdf497 5%, #fd5949 45%, #d6249f 60%, #285AEB 90%); */
}
#instagramSVG:active {
  fill: url("#instaFill");
}

/* Youtube icon colors. */
#youtubeSVG:hover,
#youtubeURL:focus #youtubeSVG {
  fill: #f00;
}
#youtubeSVG:active {
  fill: #ea0202;
}

/* Discord icon colors. */
#discordSVG:hover,
#discordURL:focus #discordSVG {
  fill: #7288d9;
}
#discordSVG:active {
  fill: #657ac4;
}

/*******************************************************************************
 *                                                                             *
 * Landing View (Bottom Styles)                                                *
 *                                                                             *
 ******************************************************************************/

/* Style for a general label on the bottom of the landing view. */
.bot_label {
  font-size: 9px;
  letter-spacing: 1px;
  font-weight: bold;
  text-shadow: 0px 0px 0px #bebcbb;
}

/* Divider used on the bottom of the landing view. */
.bot_divider {
  height: 25px;
  width: 2px;
  background: rgba(107, 105, 105, 0.7);
  margin-left: 20px;
  margin-right: 20px;
}

/* * *
* Landing View (Bottom Styles) | Left Content
* * */

/* Maintains maximum width on the status bar. */
#server_status_wrapper {
  display: inline-flex;
  width: 75px;
}

/* Span which displays the player count of the selected server. */
#player_count {
  color: #949494;
  font-size: 8px;
  font-weight: 900;
  text-shadow: 0px 0px 20px #949494;
  margin-left: 10px;
}

/* Wrapper container for the mojang status bar. */
#mojangStatusWrapper {
  position: relative;
  display: flex;
  cursor: pointer;
}

/* Icon which displays the status of the mojang services. */
#mojang_status_icon {
  font-size: 30px;
  color: #848484;
  margin-left: 15px;
  font-family: "sans-serif";
}

/* Tooltip which displays more details about the mojang statuses. */
#mojangStatusTooltip {
  position: absolute;
  visibility: hidden;
  opacity: 0;
  width: 145px;
  min-height: 150px;
  background-color: rgba(0, 0, 0, 0.75);
  color: #fff;
  border-radius: 4px;
  padding: 5px 10px;
  z-index: 1;
  font-family: "Avenir Medium";
  font-size: 12px;
  transition: visibility 0s linear 0.25s, opacity 0.25s ease;
  bottom: calc(100% + 15px);
  transform: translateX(-50%);
  margin-left: 50%;
  box-shadow: 0px 0px 20px rgb(0, 0, 0);
  cursor: default;
}
#mojangStatusTooltip:after {
  content: " ";
  position: absolute;
  left: 50%;
  top: 100%;
  margin-left: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: rgba(0, 0, 0, 0.75) transparent transparent transparent;
}
#mojangStatusWrapper:hover #mojangStatusTooltip {
  visibility: visible;
  opacity: 1;
  transition-delay: 0s;
}

/* Tooltip title for the mojang statuses. */
#mojangStatusTooltipTitle {
  width: 100%;
  text-align: center;
  margin-bottom: 5px;
  letter-spacing: 1px;
}

/* Wrapper container for the non essential services title. */
#mojangStatusNEContainer {
  display: flex;
  align-items: center;
  margin: 10px 0px;
}

/* White bar which surrounds the non essential service title. */
.mojangStatusNEBar {
  height: 1px;
  width: 100%;
  background: white;
}

/* Non essential service title text. */
#mojangStatusNETitle {
  font-size: 10px;
  padding: 0px 3px;
  text-align: center;
  letter-spacing: 1px;
}

/* Wrapper container for mojang service information. */
.mojangStatusContainer {
  display: flex;
}

/* Displays the name of the mojang service. */
.mojangStatusName {
  width: 100%;
  font-size: 10px;
  letter-spacing: 1px;
  line-height: 12px;
  padding: 6px 0px;
}

/* Displays the status of the mojang service. */
.mojangStatusIcon {
  margin-right: 10px;
  font-size: 18.5px;
  color: #848484;
}

/* * *
* Landing View (Bottom Styles) | Center Content
* * */

/* Button which opens the news view. */
#newsButton {
  background: none;
  border: none;
  cursor: pointer;
  outline: none;
}
#newsButton:hover #newsButtonText,
#newsButton:focus #newsButtonText {
  text-shadow: 0px 0px 20px #fff, 0px 0px 20px #fff;
}
#newsButton:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7, 0px 0px 20px #c7c7c7;
}

#newsButton:hover #newsButtonSVG,
#newsButton:focus #newsButtonSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #fff);
}
#newsButton:active #newsButtonSVG .arrowLine {
  stroke: #c7c7c7;
}
#newsButton:active #newsButtonSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #c7c7c7);
}
#newsButton:disabled #newsButtonSVG .arrowLine {
  stroke: rgba(255, 255, 255, 0.75);
}

/* Icon which indicates there is new news. */
#newsButtonAlert {
  width: 5px;
  height: 5px;
  position: absolute;
  border-radius: 50%;
  background: red;
  right: -1px;
  top: 50%;
}

/* Arrow image which floats above the news button. */
#newsButtonSVG {
  height: 11px;
  margin-left: -2px;
  transition: 0.25s ease;
}

/* Span which contains the news button text. */
#newsButtonText {
  color: white;
  font-weight: 900;
  letter-spacing: 2px;
  text-shadow: 0px 0px 0px #bebcbb;
  font-size: 11px;
  line-height: 30px;
  display: flex;
  transition: 0.25s ease;
}

/* * *
* Landing View (Bottom Styles) | Right Content
* * */

/* Main launch content container. */
#landingContainer > #lower > #right #launch_content {
  position: relative;
  top: 25px;
  display: inline-flex;
}

/* The launch button. */
#launch_button {
  background: none;
  border: none;
  cursor: pointer;
  font-weight: 900;
  letter-spacing: 2px;
  text-shadow: 0px 0px 0px #bebcbb;
  font-size: 20px;
  padding: 0px;
  transition: 0.25s ease;
  outline: none;
}
#launch_button:hover,
#launch_button:focus {
  text-shadow: 0px 0px 20px #fff, 0px 0px 20px #fff;
}
#launch_button:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7, 0px 0px 20px #c7c7c7;
}
#launch_button:disabled {
  color: #c7c7c7;
  cursor: default;
  pointer-events: none;
}

/* Launch details main container, hidden until launch processing begins. */
#launch_details {
  position: relative;
  top: 25px;
  display: none;
}

/* Left side of launch details container, displays percentage and a divider. */
#launch_details_left {
  display: flex;
}

/* Span which displays percentage complete. */
#launch_progress_label {
  font-weight: 900;
  letter-spacing: 1px;
  text-shadow: 0px 0px 0px #bebcbb;
  font-size: 20px;
  min-width: 53.21px;
  max-width: 53.21px;
  text-align: right;
}

/* Right side of launch details container, displays progress bar and details. */
#launch_details_right {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

/* Button which opens the server selection view. */
#server_selection_button {
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
  line-height: 24px;
  padding: 0px;
  transition: 0.25s ease;
}
#server_selection_button:hover,
#server_selection_button:focus {
  text-shadow: 0px 0px 20px #fff, 0px 0px 20px #fff, 0px 0px 20px #fff;
}
#server_selection_button:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7, 0px 0px 20px #c7c7c7, 0px 0px 20px #c7c7c7;
}

/* Progress bar styles. */
#launch_progress[value] {
  height: 3px;
  width: 265px;
  -webkit-appearance: none;
}
#launch_progress[value]::-webkit-progress-bar {
  background-color: transparent;
}
#launch_progress[value]::-webkit-progress-value {
  background-color: #fff;
}

/* Span which displays information about the status of the launch process. */
#launch_details_text {
  font-size: 11px;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}

/*******************************************************************************
 *                                                                             *
 * Overlay View (overlay.ejs)                                                  *
 *                                                                             *
 ******************************************************************************/

/* * *
* Overlay View (Main Content)
* * */

/* Overlay container, placed over the main div. */
#overlayContainer {
  position: absolute;
  z-index: 500;
  top: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: calc(100% - 22px);
  background: rgba(0, 0, 0, 0.5);
}

/* Main overlay content. */
#overlayContent {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  /*justify-content: space-between;*/
  width: 300px;
  /*height: 35%;*/
  box-sizing: border-box;
  padding: 15px 0px;
  /* background-color: #424242; */
  text-align: center;
}

/* Main overlay content anchor styles. */
#overlayContent a,
#overlayDismiss {
  color: rgba(202, 202, 202, 0.75);
  transition: 0.25s ease;
}
#overlayContent a:hover,
#overlayContent a:focus,
#overlayDismiss:focus {
  color: rgba(255, 255, 255, 0.75);
}
#overlayContent a:active,
#overlayDismiss:active {
  color: rgba(165, 165, 165, 0.75);
}

/* Add spacing between overlay content elements. */
#overlayContent > *:first-child {
  margin-top: 0px !important;
}
#overlayContent > *:last-child {
  margin-bottom: 0px !important;
}
#overlayContent > * {
  margin: 8px 0px;
}

/* Overlay title styles. */
#overlayTitle {
  font-family: "Avenir Medium";
  font-size: 20px;
  font-weight: bold;
  letter-spacing: 1px;
  -webkit-user-select: initial;
}

/* Overlay description styles. */
#overlayDesc {
  font-size: 12px;
  font-weight: bold;
  -webkit-user-select: initial;
}

/* Div which contains action buttons. */
#overlayActionContainer {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

/* Overlay acknowledge button styles. */
#overlayAcknowledge {
  background: none;
  border: 1px solid #ffffff;
  color: white;
  font-family: "Avenir Medium";
  font-weight: bold;
  border-radius: 2px;
  padding: 0px 8.1px;
  cursor: pointer;
  transition: 0.25s ease;
}
#overlayAcknowledge:hover,
#overlayAcknowledge:focus {
  box-shadow: 0px 0px 10px 0px #fff;
  outline: none;
}
#overlayAcknowledge:active {
  border-color: rgba(255, 255, 255, 0.75);
  color: rgba(255, 255, 255, 0.75);
}

/* Overlay dismiss option styles. */
#overlayDismiss {
  font-weight: bold;
  font-size: 10px;
  text-decoration: none;
  padding-top: 2.5px;
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
}
#overlayDismiss:hover {
  color: rgba(255, 255, 255, 0.75);
}
#overlayDismiss:active {
  color: rgba(165, 165, 165, 0.75);
}

/* * *
* Overlay View (Server + Account Selection Content)
* * */

/* Server selection content container. */
#serverSelectContent,
#accountSelectContent {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 75%;
}

/* Server selection header. */
#serverSelectHeader,
#accountSelectHeader {
  font-family: "Avenir Medium";
  font-size: 20px;
  font-weight: bold;
  color: #fff;
  margin-bottom: 25px;
}

/* Wrapper div for the list of available servers. */
#serverSelectList,
#accountSelectList {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  max-height: 65%;
  min-height: 40%;
}

/* Scrollable div which lists the available servers. */
#serverSelectListScrollable,
#accountSelectListScrollable {
  padding: 0px 5px;
  overflow-y: scroll;
}
#serverSelectListScrollable::-webkit-scrollbar,
#accountSelectListScrollable::-webkit-scrollbar {
  width: 2px;
}
#serverSelectListScrollable::-webkit-scrollbar-track,
#accountSelectListScrollable::-webkit-scrollbar-track {
  display: none;
}
#serverSelectListScrollable::-webkit-scrollbar-thumb,
#accountSelectListScrollable::-webkit-scrollbar-thumb {
  border-radius: 10px;
  box-shadow: inset 0 0 10px rgba(255, 255, 255, 0.5);
}

/* Content container for a server listing. */
.serverListing {
  border: none;
  padding: 0px;
  width: 375px;
  min-height: 60px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  opacity: 0.6;
  transition: 0.25s ease;
  cursor: pointer;
  position: relative;
  background: rgba(131, 131, 131, 0.25);
}
.serverListing[selected] {
  cursor: default;
  opacity: 1;
}
.serverListing:hover,
.serverListing:focus {
  outline: none;
  opacity: 1;
}

.accountListing {
  color: white;
  border: 1px solid rgba(126, 126, 126, 0.57);
  border-radius: 3px;
  padding: 5px 45px;
  width: 250px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  opacity: 0.6;
  transition: 0.25s ease;
  cursor: pointer;
  position: relative;
  background: rgba(0, 0, 0, 0.25);
}
.accountListing[selected] {
  cursor: default;
  opacity: 1;
}
.accountListing:hover,
.accountListing:focus {
  outline: none;
  opacity: 1;
}

.accountListingName {
  display: flex;
  height: 100%;
  width: 100%;
  padding-left: 10px;
}

/* Add spacing between server listings. */
#serverSelectListScrollable > .serverListing:not(:first-child):not(:last-child),
#accountSelectListScrollable
  > .accountListing:not(:first-child):not(:last-child) {
  margin: 5px 0px;
}
#serverSelectListScrollable > .serverListing:first-child,
#accountSelectListScrollable > .accountListing:first-child {
  margin-bottom: 5px;
}
#serverSelectListScrollable > .serverListing:last-child,
#accountSelectListScrollable > .accountListing:last-child {
  margin-top: 5px;
}

/* Server listing image. */
.serverListingImg {
  margin: 0px 10px 0px 5px;
  border: 1px solid #fff;
  height: 50px;
  width: 50px;
}

/* Content container for the server listing's details. */
.serverListingDetails {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: space-between;
  height: 50px;
}

/* The name of the server listing. */
.serverListingName {
  font-size: 14px;
  font-weight: bold;
}

/* Description for the server listing. */
.serverListingDescription {
  font-size: 10px;
  line-height: 10px;
  font-weight: bold;
}

/* Content container for the server listing's information. */
.serverListingInfo {
  width: 100%;
  display: flex;
  justify-content: flex-start;
}

/* The minecraft version of the server listing. */
.serverListingVersion {
  font-size: 10px;
  text-align: center;
  display: flex;
  justify-content: center;
  align-items: center;
  line-height: 12px;
  height: 12px;
  border-radius: 2px;
  background: rgba(31, 140, 11, 0.8);
  padding: 0px 2px;
}

/* The revision version of the server's manifest. */
.serverListingRevision {
  color: #969696;
  font-size: 10px;
  line-height: 12px;
  padding: 0px 5px;
}

/* Star which indicates the default (main) server. */
.serverListingStarWrapper {
  display: flex;
  align-items: center;
  cursor: pointer;
  height: 12px;
  position: relative;
}
/* Tooltip which displays when hovering over the star. */
.serverListingStarTooltip {
  visibility: hidden;
  opacity: 0;
  width: 65px;
  background-color: rgba(0, 0, 0, 0.4);
  text-align: center;
  border-radius: 4px;
  position: absolute;
  z-index: 1;
  left: 130%;
  font-size: 10px;
  transition: visibility 0s linear 0.25s, opacity 0.25s ease;
}
.serverListingStarTooltip::after {
  content: " ";
  position: absolute;
  top: 50%;
  right: 100%; /* To the left of the tooltip */
  margin-top: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: transparent rgba(0, 0, 0, 0.4) transparent transparent;
}
.serverListingStarWrapper:hover .serverListingStarTooltip {
  visibility: visible;
  opacity: 1;
  transition-delay: 0s;
}

/* Content container which contains the server select actions. */
#serverSelectActions,
#accountSelectActions {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 25px;
}

/* Server selection confirm button styles. */
#serverSelectConfirm,
#accountSelectConfirm {
  background: none;
  border: 1px solid #ffffff;
  color: white;
  font-family: "Avenir Medium";
  font-weight: bold;
  border-radius: 2px;
  padding: 0px 8.1px;
  cursor: pointer;
  transition: 0.25s ease;
  min-height: 20.67px;
}
#serverSelectConfirm:hover,
#serverSelectConfirm:focus,
#accountSelectConfirm:hover,
#accountSelectConfirm:focus {
  box-shadow: 0px 0px 10px 0px #fff;
  outline: none;
}
#serverSelectConfirm:active,
#accountSelectConfirm:active {
  border-color: rgba(255, 255, 255, 0.75);
  color: rgba(255, 255, 255, 0.75);
}

/* Server selection cancel button styles. */
#serverSelectCancel,
#accountSelectCancel {
  font-weight: bold;
  font-size: 10px;
  text-decoration: none;
  padding-top: 2.5px;
  color: rgba(202, 202, 202, 0.75);
  transition: 0.25s ease;
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
}
#serverSelectCancel:hover,
#serverSelectCancel:focus,
#accountSelectCancel:hover,
#accountSelectCancel:focus {
  color: rgba(255, 255, 255, 0.75);
}
#serverSelectCancel:active,
#accountSelectCancel:active {
  color: rgba(165, 165, 165, 0.75);
}

/*******************************************************************************
 *                                                                             *
 * Loading Element (app.ejs)                                                   *
 *                                                                             *
 ******************************************************************************/

/* Loading container, placed above everything. */
#loadingContainer {
  position: absolute;
  z-index: 400;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: calc(100% - 22px);
}

/* Loading content container. */
#loadingContent {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

/* Spinner container. */
#loadSpinnerContainer {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Stationary image for the spinner. */
#loadCenterImage {
  position: absolute;
  width: 128px;
  height: auto;
}

/* Rotating image for the spinner. */
#loadSpinnerImage {
  width: 220px;
  height: auto;
  z-index: 400;
}

/* Rotating animation for the spinner. */
@keyframes rotating {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

/* Class which is applied when the spinner image is spinning. */
.rotating {
  animation: rotating 10s linear infinite;
}

/*******************************************************************************
 *                                                                             *
 * register View (register.ejs)                                                      *
 *                                                                             *
 ******************************************************************************/

/* Styles for dimmer register span. */
.registerSpanDim {
  font-size: 12px;
  color: #848484;
  font-weight: bold;
}

/* Main register container. */
#registerContainer {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  transition: filter 0.25s ease;
  background: rgba(0, 0, 0, 0.5);
}

/* register cancel button styles. */
#registerCancelContainer {
  position: absolute;
  top: 5%;
  right: 5%;
}

/* register cancel button styles. */
#registerCancelButton {
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
  transition: 0.25s ease;
}
#registerCancelButton:hover #registerCancelIcon,
#registerCancelButton:hover #registerCancelText,
#registerCancelButton:focus #registerCancelIcon,
#registerCancelButton:focus #registerCancelText {
  text-shadow: 0px 0px 20px white;
}
#registerCancelButton:hover #registerCancelIcon,
#registerCancelButton:focus #registerCancelIcon {
  box-shadow: 0px 0px 20px white;
}
#registerCancelButton:active #registerCancelIcon,
#registerCancelButton:active #registerCancelText {
  text-shadow: 0px 0px 20px rgba(255, 255, 255, 0.75);
  color: rgba(255, 255, 255, 0.75);
  border-color: rgba(255, 255, 255, 0.75);
}
#registerCancelButton:active #registerCancelIcon {
  box-shadow: 0px 0px 20px rgba(255, 255, 255, 0.75);
}
#registerCancelButton:disabled {
  pointer-events: none;
}
#registerCancelButton:disabled #registerCancelIcon,
#registerCancelButton:disabled #registerCancelText {
  color: rgba(255, 255, 255, 0.75);
  border-color: rgba(255, 255, 255, 0.75);
}

/* The X in a circle icon for the cancel button. */
#registerCancelIcon {
  border-radius: 50%;
  border: 1px solid white;
  box-sizing: border-box;
  height: 30px;
  width: 30px;
  font-size: 19px;
  line-height: 30px;
  margin: 0 auto;
  margin-bottom: 5px;
  transition: 0.25s ease;
}
/* Text for the register cancel button. */
#registerCancelText {
  font-size: 15px;
  transition: 0.25s ease;
}

/* register content wrapper. */
#registerContent {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  padding: 0px 25px;
}

/* register form. */
#registerForm {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* register form anchor styles. */
#registerForm a {
  font-size: 12px;
  color: #848484;
  font-weight: bold;
  text-decoration: none;
  transition: 0.25s ease;
}
#registerForm a:hover,
#registerForm a:focus {
  color: #a2a2a2;
  outline: none;
}
#registerForm a:active {
  color: #8b8b8b;
}

/* Logo on register form. */
#registerImageSeal {
  border-radius: 50%;
  border: 2px solid #cad7e1;
  background: rgba(1, 2, 1, 0.5);
  height: 125px;
  width: 125px;
  box-shadow: 0px 0px 10px 0px rgb(0, 0, 0);
  margin-bottom: 20px;
}

/* Header on register view. */
#registerSubheader {
  font-family: "Avenir Medium";
  margin-bottom: 25px;
  font-size: 12px;
  letter-spacing: 1px;
  font-weight: bold;
}

/* Container to organize register field elements. */
.registerFieldContainer {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* SVG icons on the register view. */
.registerSVG {
  fill: #fff;
  height: 20px;
  width: 20px;
}

/* Span which displays errors related to register field content. */
.registerErrorSpan {
  font-family: "Avenir Medium";
  font-weight: bold;
  font-size: 0.7em;
  color: #ff1b0c;
  width: 100%;
  text-align: right;
  position: absolute;
  top: 7px;
  opacity: 0;
  transition: 0.25s ease;
}

/* register text input styles. */
.registerField {
  font-family: "Avenir Book";
  background: none;
  border-width: 1.5px 0px 0px 0px;
  border-style: solid;
  width: 250px;
  margin-bottom: 20px;
  border-color: #fff;
  color: rgba(255, 255, 255, 0.75);
  font-weight: bold;
  text-align: center;
  box-sizing: border-box;
  padding: 7.5px;
  font-size: 10px;
  letter-spacing: 1px;
}
.registerField:focus {
  outline: none;
}
.registerField:disabled {
  color: rgba(255, 255, 255, 0.5);
}
.registerField::-webkit-input-placeholder {
  color: rgba(255, 255, 255, 0.75);
  font-size: 10px;
  letter-spacing: 1px;
  text-align: center;
  font-weight: bold;
}
.registerField:focus::-webkit-input-placeholder {
  color: transparent;
}

/* Container which contains the forgot and remember options. */
#registerOptions {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

/* Remember option text. */
#registerRememberText {
  padding-right: 10px;
  transition: 0.25s ease;
}

/* register button styles. */
#registerButton {
  background: none;
  font-weight: bold;
  letter-spacing: 2px;
  border: none;
  padding: 15px 5px;
  margin: 10px 0px;
  cursor: pointer;
  position: relative;
  right: -20px;
  transition: 0.5s ease;
}
#registerButton:disabled {
  color: rgba(255, 255, 255, 0.75);
  pointer-events: none;
}
#registerButton[loading] {
  color: #fff;
}
#registerButton:hover,
#registerButton:focus {
  text-shadow: 0px 0px 20px #fff;
  outline: none;
}
#registerButton:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7;
}
#registerSVG {
  -webkit-transform: translate3d(0, 0, 0);
  overflow: visible;
  transform: rotate(90deg);
  margin-left: 20px;
  transition: 0.25s ease;
  width: 20px;
  height: 20px;
}
#registerButton:hover #registerSVG,
#registerButton:focus #registerSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #fff);
}
#registerButton:active #registerSVG .arrowLine {
  stroke: #c7c7c7;
}
#registerButton:active #registerSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #c7c7c7);
}
#registerButton:disabled #registerSVG .arrowLine {
  stroke: rgba(255, 255, 255, 0.75);
}

#registerButtonContent {
  display: flex;
  align-items: center;
}

#registerButton .circle-loader,
#registerButton[loading] #registerSVG {
  display: none;
}
#registerButton[loading] .circle-loader,
#registerButton #registerSVG {
  display: initial;
}

#cancelRegister {
  background: none;
  border: none;
  margin-top: 20px;
  font-size: 0.8em;
}

/* Disclaimer container. */
#registerDisclaimer {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* Add spacing between register anchor and disclaimer. */
#registerRegisterSpan {
  margin-bottom: 5px;
}

/* Disclaimer text styles. */
.registerDisclaimerText {
  font-size: 7px;
  color: #848484;
  font-weight: bold;
  text-align: center;
}

/* * *
* register View | Custom Checkbox
* * */

/* Checkbox container. */
#checkmarkContainer {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  position: relative;
  cursor: pointer;
  font-size: 22px;
  -webkit-user-select: none;
}

/* Hide the default checkbox. */
#checkmarkContainer input {
  opacity: 0;
  cursor: pointer;
  position: absolute;
}

/* Create a custom checkbox. */
.registerCheckmark {
  position: relative;
  height: 10px;
  width: 10px;
  border: 1px solid #848484;
  border-radius: 1px;
  background: none;
  transition: 0.25s ease;
}
/* On hover and focus, add a grey border color. */
#checkmarkContainer:hover input ~ *,
#checkmarkContainer input:focus ~ * {
  color: #a2a2a2;
  border-color: #a2a2a2;
}
/* On keydown, darken the checkbox a bit. */
#checkmarkContainer input:active ~ *:not(#registerRememberText) {
  color: #8d8d8d;
  border-color: #8d8d8d;
}
#checkmarkContainer[disabled] {
  pointer-events: none;
}
/* For checked -> #checkmarkContainer input:checked ~ * */
/* Create the checkmark/indicator (hidden when not checked). */
.registerCheckmark:after {
  content: "";
  display: none;
}
/* Show the checkmark when checked. */
#checkmarkContainer input:checked ~ .registerCheckmark:after {
  display: block;
}
/* Style the checkmark/indicator. */
#checkmarkContainer .registerCheckmark:after {
  position: absolute;
  left: 3.5px;
  top: 0.5px;
  width: 2px;
  height: 6px;
  border: solid #a2a2a2;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

/*
#register_filter {
    height: calc(100% - 22px);
    width: 100%;
    z-index: 9000;
    position: absolute;
    filter: blur(8px) contrast(0.9) brightness(1.0);
    background: url('./../images/backgrounds/0.jpg') no-repeat center center fixed;
    transform: scale(1.2);
    background-size: cover;
}
*/

/*******************************************************************************
 *                                                                             *
 * auth View (auth.ejs)                                                      *
 *                                                                             *
 ******************************************************************************/

/* Styles for dimmer auth span. */
.authSpanDim {
  font-size: 12px;
  color: #848484;
  font-weight: bold;
}

/* Main auth container. */
#authContainer {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  transition: filter 0.25s ease;
  background: rgba(0, 0, 0, 0.5);
}

/* auth cancel button styles. */
#authCancelContainer {
  position: absolute;
  top: 5%;
  right: 5%;
}

/* auth cancel button styles. */
#authCancelButton {
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
  transition: 0.25s ease;
}
#authCancelButton:hover #authCancelIcon,
#authCancelButton:hover #authCancelText,
#authCancelButton:focus #authCancelIcon,
#authCancelButton:focus #authCancelText {
  text-shadow: 0px 0px 20px white;
}
#authCancelButton:hover #authCancelIcon,
#authCancelButton:focus #authCancelIcon {
  box-shadow: 0px 0px 20px white;
}
#authCancelButton:active #authCancelIcon,
#authCancelButton:active #authCancelText {
  text-shadow: 0px 0px 20px rgba(255, 255, 255, 0.75);
  color: rgba(255, 255, 255, 0.75);
  border-color: rgba(255, 255, 255, 0.75);
}
#authCancelButton:active #authCancelIcon {
  box-shadow: 0px 0px 20px rgba(255, 255, 255, 0.75);
}
#authCancelButton:disabled {
  pointer-events: none;
}
#authCancelButton:disabled #authCancelIcon,
#authCancelButton:disabled #authCancelText {
  color: rgba(255, 255, 255, 0.75);
  border-color: rgba(255, 255, 255, 0.75);
}

/* The X in a circle icon for the cancel button. */
#authCancelIcon {
  border-radius: 50%;
  border: 1px solid white;
  box-sizing: border-box;
  height: 30px;
  width: 30px;
  font-size: 19px;
  line-height: 30px;
  margin: 0 auto;
  margin-bottom: 5px;
  transition: 0.25s ease;
}
/* Text for the auth cancel button. */
#authCancelText {
  font-size: 15px;
  transition: 0.25s ease;
}

/* auth content wrapper. */
#authContent {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  padding: 0px 25px;
}

/* auth form. */
#authForm {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* auth form anchor styles. */
#authForm a {
  font-size: 12px;
  color: #848484;
  font-weight: bold;
  text-decoration: none;
  transition: 0.25s ease;
}
#authForm a:hover,
#authForm a:focus {
  color: #a2a2a2;
  outline: none;
}
#authForm a:active {
  color: #8b8b8b;
}

/* Logo on auth form. */
#authImageSeal {
  border-radius: 50%;
  border: 2px solid #cad7e1;
  background: rgba(1, 2, 1, 0.5);
  height: 125px;
  width: 125px;
  box-shadow: 0px 0px 10px 0px rgb(0, 0, 0);
  margin-bottom: 20px;
}

/* Header on auth view. */
#authSubheader {
  font-family: "Avenir Medium";
  margin-bottom: 25px;
  font-size: 12px;
  letter-spacing: 1px;
  font-weight: bold;
}

/* Container to organize auth field elements. */
.authFieldContainer {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* SVG icons on the auth view. */
.authSVG {
  fill: #fff;
  height: 20px;
  width: 20px;
}

/* Span which displays errors related to auth field content. */
.authErrorSpan {
  font-family: "Avenir Medium";
  font-weight: bold;
  font-size: 0.7em;
  color: #ff1b0c;
  width: 100%;
  text-align: right;
  position: absolute;
  top: 7px;
  opacity: 0;
  transition: 0.25s ease;
}

/* auth text input styles. */
.authField {
  font-family: "Avenir Book";
  background: none;
  border-width: 1.5px 0px 0px 0px;
  border-style: solid;
  width: 250px;
  margin-bottom: 20px;
  border-color: #fff;
  color: rgba(255, 255, 255, 0.75);
  font-weight: bold;
  text-align: center;
  box-sizing: border-box;
  padding: 7.5px;
  font-size: 10px;
  letter-spacing: 1px;
}
.authField:focus {
  outline: none;
}
.authField:disabled {
  color: rgba(255, 255, 255, 0.5);
}
.authField::-webkit-input-placeholder {
  color: rgba(255, 255, 255, 0.75);
  font-size: 10px;
  letter-spacing: 1px;
  text-align: center;
  font-weight: bold;
}
.authField:focus::-webkit-input-placeholder {
  color: transparent;
}

/* Container which contains the forgot and remember options. */
#authOptions {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

/* Remember option text. */
#authRememberText {
  padding-right: 10px;
  transition: 0.25s ease;
}

/* auth button styles. */
#authButton {
  background: none;
  font-weight: bold;
  letter-spacing: 2px;
  border: none;
  padding: 15px 5px;
  margin: 10px 0px;
  cursor: pointer;
  position: relative;
  right: -20px;
  transition: 0.5s ease;
}
#authButton:disabled {
  color: rgba(255, 255, 255, 0.75);
  pointer-events: none;
}
#authButton[loading] {
  color: #fff;
}
#authButton:hover,
#authButton:focus {
  text-shadow: 0px 0px 20px #fff;
  outline: none;
}
#authButton:active {
  color: #c7c7c7;
  text-shadow: 0px 0px 20px #c7c7c7;
}
#authSVG {
  -webkit-transform: translate3d(0, 0, 0);
  overflow: visible;
  transform: rotate(90deg);
  margin-left: 20px;
  transition: 0.25s ease;
  width: 20px;
  height: 20px;
}
#authButton:hover #authSVG,
#authButton:focus #authSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #fff);
}
#authButton:active #authSVG .arrowLine {
  stroke: #c7c7c7;
}
#authButton:active #authSVG {
  -webkit-filter: drop-shadow(0px 0px 2px #c7c7c7);
}
#authButton:disabled #authSVG .arrowLine {
  stroke: rgba(255, 255, 255, 0.75);
}

#authButtonContent {
  display: flex;
  align-items: center;
}

#authButton .circle-loader,
#authButton[loading] #authSVG {
  display: none;
}
#authButton[loading] .circle-loader,
#authButton #authSVG {
  display: initial;
}

/* Disclaimer container. */
#authDisclaimer {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* Add spacing between auth anchor and disclaimer. */
#authauthSpan {
  margin-bottom: 5px;
}

/* Disclaimer text styles. */
.authDisclaimerText {
  font-size: 7px;
  color: #848484;
  font-weight: bold;
  text-align: center;
}

/* * *
* auth View | Custom Checkbox
* * */

/* Checkbox container. */
#checkmarkContainer {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  position: relative;
  cursor: pointer;
  font-size: 22px;
  -webkit-user-select: none;
}

/* Hide the default checkbox. */
#checkmarkContainer input {
  opacity: 0;
  cursor: pointer;
  position: absolute;
}

/* Create a custom checkbox. */
.authCheckmark {
  position: relative;
  height: 10px;
  width: 10px;
  border: 1px solid #848484;
  border-radius: 1px;
  background: none;
  transition: 0.25s ease;
}
/* On hover and focus, add a grey border color. */
#checkmarkContainer:hover input ~ *,
#checkmarkContainer input:focus ~ * {
  color: #a2a2a2;
  border-color: #a2a2a2;
}
/* On keydown, darken the checkbox a bit. */
#checkmarkContainer input:active ~ *:not(#authRememberText) {
  color: #8d8d8d;
  border-color: #8d8d8d;
}
#checkmarkContainer[disabled] {
  pointer-events: none;
}
/* For checked -> #checkmarkContainer input:checked ~ * */
/* Create the checkmark/indicator (hidden when not checked). */
.authCheckmark:after {
  content: "";
  display: none;
}
/* Show the checkmark when checked. */
#checkmarkContainer input:checked ~ .authCheckmark:after {
  display: block;
}
/* Style the checkmark/indicator. */
#checkmarkContainer .authCheckmark:after {
  position: absolute;
  left: 3.5px;
  top: 0.5px;
  width: 2px;
  height: 6px;
  border: solid #a2a2a2;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.passwordVisibleImg {
  width: 20px;
  height: 20px;
  background: url("/../images/closed-eye.png") no-repeat;
  background-size: contain;
  background-position: center;
  position: absolute;
  right: 0;
}

.passwordVisible {
  background: url("/../images/open-eye.png") no-repeat;
  background-size: contain;
  background-position: center;
}

/*
#auth_filter {
    height: calc(100% - 22px);
    width: 100%;
    z-index: 9000;
    position: absolute;
    filter: blur(8px) contrast(0.9) brightness(1.0);
    background: url('./../images/backgrounds/0.jpg') no-repeat center center fixed;
    transform: scale(1.2);
    background-size: cover;
}
*/

/*******************************************************************************
 *                                                                             *
 * rules View (rules.ejs)                                                  *
 *                                                                             *
 ******************************************************************************/

/* * *
* rules View (Main Content)
* * */

/* rules container, placed over the main div. */
#rulesContainer {
  position: relative;
  z-index: 500;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
}

/* Main rules content. */
#rulesContent {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-sizing: border-box;
  /* background-color: #424242; */
  text-align: left;
  border-radius: 5px;
  height: 80%;
  margin-left: 100px;
  margin-right: 100px;
  margin-bottom: 30px;
}

#rulesContentScrollable {
  overflow: scroll;
  overflow-x: hidden;
  font-family: "Avenir Next Cyr";
}
#rulesContentScrollable span {
  display: block;
}

#rulesContentScrollable::-webkit-scrollbar-thumb {
  border-radius: 10px;
  box-shadow: inset 0 0 10px rgba(255, 255, 255, 0.5);
}
#rulesContentScrollable::-webkit-scrollbar-track {
  display: none;
}
#rulesContentScrollable::-webkit-scrollbar {
  width: 2px;
}

/* Main rules content anchor styles. */
#rulesContent a,
#rulesDismiss {
  color: rgba(202, 202, 202, 0.75);
  transition: 0.25s ease;
}
#rulesContent a:hover,
#rulesContent a:focus,
#rulesDismiss:focus {
  color: rgba(255, 255, 255, 0.75);
}
#rulesContent a:active,
#rulesDismiss:active {
  color: rgba(165, 165, 165, 0.75);
}

/* Add spacing between rules content elements. */
#rulesContent > *:first-child {
  margin-top: 0px !important;
}
#rulesContent > *:last-child {
  margin-bottom: 0px !important;
}
#rulesContent > * {
  margin: 8px 0px;
}

/* rules title styles. */
#rulesTitle {
  font-family: "Avenir Medium";
  font-size: 20px;
  font-weight: bold;
  letter-spacing: 1px;
  -webkit-user-select: initial;
}

/* rules description styles. */
#rulesDesc {
  font-size: 12px;
  font-weight: bold;
  -webkit-user-select: initial;
}

/* Div which contains action buttons. */
#rulesActionContainer {
  display: flex;
  align-items: flex-end;
  margin-top: 10px;
}

/* rules acknowledge button styles. */
#rulesAcknowledge {
  background: none;
  border: 1px solid #ffffff;
  color: white;
  font-family: "Avenir Medium";
  font-weight: bold;
  border-radius: 2px;
  padding: 0px 8.1px;
  cursor: pointer;
  transition: 0.25s ease;
  margin-left: 5px;
}
#rulesAcknowledge:hover,
#rulesAcknowledge:focus {
  box-shadow: 0px 0px 10px 0px #fff;
  outline: none;
}
#rulesAcknowledge:active {
  border-color: rgba(255, 255, 255, 0.75);
  color: rgba(255, 255, 255, 0.75);
}

/* rules dismiss option styles. */
#rulesDismiss {
  font-weight: bold;
  font-size: 10px;
  text-decoration: none;
  padding-top: 2.5px;
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
}
#rulesDismiss:hover {
  color: rgba(255, 255, 255, 0.75);
}
#rulesDismiss:active {
  color: rgba(165, 165, 165, 0.75);
}

/*******************************************************************************
 *                                                                             *
 * agreement View (agreement.ejs)                                                  *
 *                                                                             *
 ******************************************************************************/

/* * *
* agreement View (Main Content)
* * */

/* agreement container, placed over the main div. */
#agreementContainer {
  position: relative;
  z-index: 500;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
}

/* Main agreement content. */
#agreementContent {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-sizing: border-box;
  padding: 15px 0px;
  /* background-color: #424242; */
  text-align: left;
  border-radius: 5px;
  /* border: solid white 2px; */
  height: 80%;
  margin-left: 100px;
  margin-right: 100px;
  margin-bottom: 30px;
}

#agreementContentScrollable {
  overflow: scroll;
  overflow-x: hidden;
  font-family: "Avenir Next Cyr";
}

#agreementContent span {
  display: block;
}

#agreementContentScrollable::-webkit-scrollbar-thumb {
  border-radius: 10px;
  box-shadow: inset 0 0 10px rgba(255, 255, 255, 0.5);
}
#agreementContentScrollable::-webkit-scrollbar-track {
  display: none;
}
#agreementContentScrollable::-webkit-scrollbar {
  width: 2px;
}

/* Main agreement content anchor styles. */
#agreementContent a,
#agreementDismiss {
  color: rgba(202, 202, 202, 0.75);
  transition: 0.25s ease;
}
#agreementContent a:hover,
#agreementContent a:focus,
#agreementDismiss:focus {
  color: rgba(255, 255, 255, 0.75);
}
#agreementContent a:active,
#agreementDismiss:active {
  color: rgba(165, 165, 165, 0.75);
}

/* Add spacing between agreement content elements. */
#agreementContent > *:first-child {
  margin-top: 0px !important;
}
#agreementContent > *:last-child {
  margin-bottom: 0px !important;
}
#agreementContent > * {
  margin: 8px 0px;
}

/* agreement title styles. */
#agreementTitle {
  font-family: "Avenir Medium";
  font-size: 20px;
  font-weight: bold;
  letter-spacing: 1px;
  -webkit-user-select: initial;
}

/* agreement description styles. */
#agreementDesc {
  font-size: 12px;
  font-weight: bold;
  -webkit-user-select: initial;
}

/* Div which contains action buttons. */
#agreementActionContainer {
  display: flex;
  align-items: flex-end;
  margin-top: 10px;
}

/* agreement acknowledge button styles. */
#agreementAcknowledge {
  background: none;
  border: 1px solid #ffffff;
  color: white;
  font-family: "Avenir Medium";
  font-weight: bold;
  border-radius: 2px;
  padding: 0px 8.1px;
  cursor: pointer;
  transition: 0.25s ease;
  margin-left: 5px;
}
#agreementAcknowledge:hover,
#agreementAcknowledge:focus {
  box-shadow: 0px 0px 10px 0px #fff;
  outline: none;
}
#agreementAcknowledge:active {
  border-color: rgba(255, 255, 255, 0.75);
  color: rgba(255, 255, 255, 0.75);
}

/* agreement dismiss option styles. */
#agreementDismiss {
  font-weight: bold;
  font-size: 10px;
  text-decoration: none;
  padding-top: 2.5px;
  background: none;
  border: none;
  outline: none;
  cursor: pointer;
}
#agreementDismiss:hover {
  color: rgba(255, 255, 255, 0.75);
}
#agreementDismiss:active {
  color: rgba(165, 165, 165, 0.75);
}
</style>