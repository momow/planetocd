class LikesController < BaseController
    def index
      if params[:password] != ENV['PLANETOCD_ADMIN_PASSWORD']
        return not_found
      end
      @likes = Like.all().order(created_at: :desc)
    end

    def create
      Like.create(like_params)
      @likes_count = article.likes_count
    end

    private

    def like_params
      params
        .permit(:article_id, :language)
        .merge(user_agent: request.user_agent,
               ip_address: request.remote_ip)
    end

    def article
      Rails.application.articles[I18n.locale.to_s][params[:article_id].to_i]
    end
end
