class AddLanguageToLikes < ActiveRecord::Migration[5.2]
  def change
    add_column :likes, :language, :string
  end
end
