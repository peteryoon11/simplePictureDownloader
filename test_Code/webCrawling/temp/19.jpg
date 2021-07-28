;(function($) {

	var Area = {};

	Area.Skin = (function() {
		var $body = $(document.body),
			$areaSkin = $(".wrap_skin");

		var openMenu = function() {
			$body.addClass("layer_on");
		};

		var closeMenu = function() {
			$body.removeClass("layer_on");
		};

		var init = function() {
			$areaSkin.on("click", ".btn_menu", openMenu);
			$areaSkin.on("click", ".btn_close", closeMenu);
		};

		return {
			init: init
		}
	})();

	Area.Profile = (function() {
		var $areaProfile = $(".area_profile");

		var toggleProfileMenu = function() {
			$areaProfile.toggleClass("on");
		};

		var init = function() {
			$areaProfile.on("click", ".btn_name", toggleProfileMenu);
		};

		return {
			init: init
		}
	})();

	Area.Category = (function() {
		var $areaNavi = $(".area_navi");

		var toggleCategory = function() {
			$areaNavi.toggleClass("on");
		};

		var init = function() {
			$areaNavi.on("click", ".btn_cate", toggleCategory);
		};

		return {
			init: init
		}
	})();


	Area.Search = (function() {
		var $areaSearch = $(".area_search"),
			$input = $areaSearch.find(".tf_search");

		var openSearch = function() {
			$areaSearch.addClass("on");
			$input.focus();
		};

		var leaveSearch = function() {
			if ($input.val() == "") {
				$areaSearch.removeClass("on");
			}
		};

		var init = function() {
			$areaSearch.on("click", ".btn_search", openSearch);
			$input.on("blur", leaveSearch);
		};

		return {
			init: init
		}
	})();

	Area.Comment = (function() {
		var $btnOpen = $(".btn_reply"),
			$fieldReply = $(".fld_reply");

		var changeStatus = function() {
			$btnOpen.toggleClass("on");
		};

		var init = function() {
			if ($fieldReply.is(":visible")) {
				$btnOpen.addClass("on");
			}
		};

		return {
			init: init,
			changeStatus: changeStatus
		}
	})();

	Area.SocialShare = (function() {

		var SOCIAL_SHARE = {
			kakaostory: {
				url: "https://story.kakao.com/share",
				width: 640,
				height: 400,
				makeShareUrl: function(url) {
					return this.url + "?url=" + encodeURIComponent(url);
				}
			},
			facebook: {
				url: "https://www.facebook.com/sharer/sharer.php",
				width: 640,
				height: 400,
				makeShareUrl: function(url) {
					return this.url + "?display=popup&u=" + encodeURIComponent(url);
				}
			},
			twitter: {
				url: "https://twitter.com/intent/tweet",
				width: 640,
				height: 400,
				makeShareUrl: function(url) {
					return this.url + "?url=" + encodeURIComponent(url);
				}
			}
		};


		var onClick = function(e) {
			var $this = $(this),
				service = SOCIAL_SHARE[$this.data("service")],
				url = location.href;

			if (service) {
				e.preventDefault();
				window.open(service.makeShareUrl(url), "", "width=" + service.width + ", height=" + service.height);
			}
		};


		var init = function() {
			$(".list_share").on("click", "a", onClick);
		};

		return {
			init: init
		}
	})();

	Area.init = function() {
		Area.Skin.init();
		Area.Profile.init();
		Area.Category.init();
		Area.Search.init();
		Area.Comment.init();
		Area.SocialShare.init();
	};

	$.Area = Area;

	var $allVideos = $("iframe[src^='//player.vimeo.com'], iframe[src^='//www.youtube.com'], object, embed, iframe[src^='http://www.youtube.com'], iframe[src^='https://www.youtube.com'], iframe[src^='http://videofarm.daum.net'], iframe[src^='https://videofarm.daum.net'], iframe[src^='//videofarm.daum.net']"),
		$fluidEl = $(".area_view");

	$allVideos.each(function() {

		$(this)
		// jQuery .data does not work on object/embed elements
			.attr('data-aspectRatio', this.height / this.width)
			.removeAttr('height')
			.removeAttr('width');

	});

	$(window).resize(function() {

		var newWidth = $fluidEl.width();
		$allVideos.each(function() {

			var $el = $(this);
			$el
				.width(newWidth)
				.height(newWidth * $el.attr('data-aspectRatio'));

		});

	}).resize();

})(jQuery);
