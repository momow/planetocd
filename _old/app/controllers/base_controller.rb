class BaseController < ActionController::Base
    layout 'default'
    before_action :redirect_domain
    before_action :set_locale

    @@redirects = {
        "planetetoc.org" => 302,
        "planetetoc.fr" => 302,
        "www.planetocd.org" => 302,
        "www.planetetoc.fr" => 302,
        "www.planetetoc.org" => 302
    }

    def redirect_domain
        status = @@redirects[request.host]
        if !status.nil?
            path = request.fullpath
            if path[0..2] == '/fr'
                path = path[3..-1]
            end
            redirect_to "#{request.protocol}#{Constants::DOMAIN}#{path}", :status => status
        end
    end

    def set_locale
        I18n.locale = "fr"
    end

    def not_found
        render 'errors/not_found', :status => 404
    end
end
